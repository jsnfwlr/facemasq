package portscan

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"facemasq/lib/db"
)

func DiscoverScanAndStoreAsync(scanID int64) (err error) {
	var devices []DeviceToScan

	if PortScan {
		devices, err = Discover()

		for _, scan := range devices {
			result := ScanDeviceAsync(scan.IPv4)
			for _, result := range result.Ports {
				if result.State == "open" {
					sql := `INSERT INTO Ports (AddressID, ScanID, Port, Protocol) VALUES (?,?,?,?);`
					_, err = db.Conn.Exec(sql, scan.AddressID, scanID, result.Number, result.Protocol)
					if err != nil {
						log.Printf("could not record port state: %v\n", err)
						return
					}
				}
			}
		}
	}
	return
}

func Discover() (devices []DeviceToScan, err error) {
	sql := `SELECT Addresses.ID AS AddressID, IPv4 FROM Addresses JOIN Interfaces ON Interfaces.ID = InterfaceID JOIN Devices ON Devices.ID = DeviceID WHERE Devices.StatusID = 1 AND Devices.IsOnline = true;`
	err = db.Conn.NewRaw(sql).Scan(db.Context, &devices)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			err = fmt.Errorf("could not find devices to scan: %v", err)
			return
		}
	}

	return
}

func ScanPortAsync(async chan Port, protocol, ipv4 string, portNum int) {
	port := Port{
		Number:   portNum,
		Protocol: protocol,
		State:    "open",
	}
	address := ipv4 + ":" + strconv.Itoa(portNum)
	_, err := net.DialTimeout(protocol, address, time.Duration(timeOut)*time.Millisecond)
	if err != nil {
		if err.Error() == fmt.Sprintf("dial %s %s:%d: connect: connection refused", protocol, ipv4, portNum) {
			port.State = "closed"
		} else if err.Error() == fmt.Sprintf("dial %s %s:%d: connect: no route to host", protocol, ipv4, portNum) {
			port.State = "unavailable"
		} else {
			port.State = "filtered"
		}
	}

	async <- port
}

func ScanDeviceAsync(ipv4 string) (scan ResultSet) {
	protocols := []string{"tcp", "udp"}
	async := make(chan Port, len(portList)*len(protocols))
	defer close(async)

	scan.Address = ipv4
	for _, protocol := range protocols {
		for portNum := range portList {
			go ScanPortAsync(async, protocol, ipv4, portNum)
			port := <-async
			scan.Ports = append(scan.Ports, port)
		}
	}
	return scan
}

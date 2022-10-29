package port

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"facemasq/lib/db"
	"facemasq/lib/logging"
	"facemasq/lib/scans"
	"facemasq/models"
)

var (
	scanWidth string
	timeOut   int
	portList  []int64
	PortScan  bool
)

func init() {
	var err error
	timeOut, err = strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		timeOut = 2000
	}

	scanWidth = os.Getenv("PORTSCAN_WIDTH")
	if scanWidth == "" || scanWidth == "narrow" {
		for i := int64(1); i <= 1024; i++ {
			portList = append(portList, i)
		}
	} else if scanWidth == "wide" {
		for i := int64(1); i <= 10000; i++ {
			portList = append(portList, i)
		}
	} else if scanWidth == "all" {
		for i := int64(1); i <= 65535; i++ {
			portList = append(portList, i)
		}
	}

	PortScan = false
	if strings.ToLower(os.Getenv("PORTSCAN")) == "true" || strings.ToLower(os.Getenv("PORTSCAN")) == "t" || os.Getenv("PORTSCAN") == "1" {
		PortScan = true
	}

	logging.Processf("%v timeout", time.Duration(timeOut)*time.Millisecond)
}

func Scan(scanID int64, scanAll bool) (err error) {
	var addresses []scans.AddressToPortScan

	if PortScan {
		addresses, err = Discover(scanAll)
		if err != nil {
			return
		}
		var ports []models.Port
		for _, scan := range addresses {
			logging.Printf(2, "Scanning %s", scan.IPv4)
			result := ScanAddress(scan.IPv4)
			for _, result := range result.Ports {
				ports = append(ports, models.Port{AddressID: int64(scan.AddressID), ScanID: scanID, Port: result.Number, Protocol: result.Protocol})
			}
		}
		if len(ports) > 0 {
			_, err = db.Conn.NewInsert().Model(&ports).Exec(db.Context)
			if err != nil {
				logging.Errorf("could not record port state: %v", err)
				return
			}
		}
	}
	return
}

func Discover(scanAll bool) (devices []scans.AddressToPortScan, err error) {
	var sql string
	sql = `SELECT addresses.id AS address_id, ipv4 FROM addresses JOIN interfaces ON interfaces.id = addresses.interface_id JOIN devices ON devices.id = interfaces.device_id WHERE devices.status_id = 1 AND devices.is_online = true;`
	if scanAll {
		sql = `SELECT addresses.id AS address_id, ipv4 FROM addresses JOIN interfaces ON interfaces.id = addresses.interface_id JOIN devices ON devices.id = interfaces.device_id WHERE devices.is_online = true;`
	}
	logging.Println(2, sql)
	err = db.Conn.NewRaw(sql).Scan(db.Context, &devices)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			err = fmt.Errorf("could not find devices to scan: %v", err)
			return
		}
	}
	if scanAll && len(devices) > 5 {
		devices = devices[0:5]
	}

	return
}

func ScanPort(protocol, ipv4 string, portNum int64) portDetails {
	address := fmt.Sprintf("%s:%d", ipv4, portNum)
	logging.Printf(2, "Scanning %s", address)
	conn, err := net.DialTimeout(protocol, address, time.Duration(timeOut)*time.Millisecond)
	port := portDetails{
		Number:   portNum,
		Protocol: protocol,
		State:    "open",
	}
	if err != nil {
		if err.Error() == fmt.Sprintf("dial %s %s:%d: connect: connection refused", protocol, ipv4, portNum) {
			port.State = "closed"
		} else if err.Error() == fmt.Sprintf("dial %s %s:%d: connect: no route to host", protocol, ipv4, portNum) {
			port.State = "unavailable"
		} else {
			port.State = "filtered"
		}
	}
	if conn == nil {
		port.State = "closed"
	}
	return port
}

func ScanAddress(ipv4 string) (scan addressPortRecord) {
	// protocols := []string{"tcp", "udp"}
	protocols := []string{"tcp"}
	scan.Address = ipv4

	for _, protocol := range protocols {
		for _, portNum := range portList {
			port := ScanPort(protocol, ipv4, portNum)
			if port.State == "open" {
				logging.Printf(0, "%s has port %d open", ipv4, portNum)
				scan.Ports = append(scan.Ports, port)
			}
		}
	}
	return
}

package port

import (
	"fmt"
	"net"
	"time"

	"facemasq/lib/db"
	"facemasq/lib/events"
	"facemasq/lib/logging"
	"facemasq/lib/scans"
	"facemasq/models"
)

func ScanAysnc(scanID int64, scanAll bool) (err error) {
	var addresses []scans.AddressToPortScan

	if PortScan {
		addresses, err = Discover(scanAll)
		var ports []models.Port
		for _, scan := range addresses {
			result := ScanAddressAsync(scan.IPv4)
			for _, result := range result.Ports {
				ports = append(ports, models.Port{AddressID: int64(scan.AddressID), ScanID: scanID, Port: result.Number, Protocol: result.Protocol})
			}
		}
		if len(ports) > 0 {
			_, err = db.Conn.NewInsert().Model(&ports).Exec(db.Context)
			if err != nil {
				logging.Error("could not record port state: %v", err)
				return
			}
		}
		events.Emit(events.Event{Kind: "scan:port:complete"})
	}
	return
}

func ScanPortAsync(async chan portDetails, protocol, ipv4 string, portNum int64) {
	port := portDetails{
		Number:   portNum,
		Protocol: protocol,
		State:    "open",
	}
	// 	testPacket := []byte("xyz")
	address := fmt.Sprintf("%s:%d", ipv4, portNum)
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
	// var udpResp int
	// if protocol == "udp" {
	// 	if conn != nil {
	// 		udpResp, err = conn.Write(testPacket)
	// 		if err != nil {
	// 			port.State = "closed"
	// 		}
	// 		fmt.Println(udpResp)

	// 	} else {
	// 		port.State = "closed"
	// 	}
	// }

	async <- port
}

func ScanAddressAsync(ipv4 string) (scan addressPortRecord) {
	var portNum int64
	// protocols := []string{"tcp", "udp"}
	protocols := []string{"tcp"}
	async := make(chan portDetails, len(portList)*len(protocols))
	defer close(async)

	scan.Address = ipv4
	for _, protocol := range protocols {
		for _, portNum = range portList {
			go ScanPortAsync(async, protocol, ipv4, portNum)
			port := <-async
			if port.State == "open" {
				scan.Ports = append(scan.Ports, port)
			}
		}
	}
	return scan
}

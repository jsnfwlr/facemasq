package portscan

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

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

func ScanAsync(ipv4 string) (scan ResultSet) {
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

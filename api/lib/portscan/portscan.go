package portscan

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var scanWidth string
var timeOut int
var portList []int
var PortScan bool

func init() {
	var err error
	timeOut, err = strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		timeOut = 2000
	}

	scanWidth = os.Getenv("SCANWIDTH")
	if scanWidth == "" {
		for i := 1; i <= 1024; i++ {
			portList = append(portList, i)
		}
	} else if scanWidth == "wide" {
		for i := 1; i <= 65535; i++ {
			portList = append(portList, i)
		}

	}

	PortScan = false
	if strings.ToLower(os.Getenv("PORTSCAN")) == "true" || strings.ToLower(os.Getenv("PORTSCAN")) == "t" || os.Getenv("PORTSCAN") == "1" {
		PortScan = true
	}

	log.Printf("%v timeout\n", time.Duration(timeOut)*time.Millisecond)
}

func ScanPort(protocol, ipv4 string, portNum int) Port {
	address := ipv4 + ":" + strconv.Itoa(portNum)
	conn, err := net.DialTimeout(protocol, address, time.Duration(timeOut)*time.Millisecond)
	port := Port{
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
	defer conn.Close()
	return port
}

func Scan(ipv4 string) (scan ResultSet) {
	protocols := []string{"tcp", "udp"}
	scan.Address = ipv4

	for _, protocol := range protocols {
		for portNum := range portList {
			port := ScanPort(protocol, ipv4, portNum)
			scan.Ports = append(scan.Ports, port)
		}
	}
	return
}

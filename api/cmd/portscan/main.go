package main

import (
	"fmt"

	"facemasq/lib/portscan"
)

func main() {
	ipv4 := "192.168.0.24"
	scan := portscan.ScanDeviceAsync(ipv4)
	fmt.Printf("%s Open ports:\n", ipv4)
	for _, port := range scan.Ports {
		if port.State == "open" && port.Protocol == "tcp" {
			fmt.Printf("%d (%s)\n", port.Number, port.Protocol)
		}
	}
}

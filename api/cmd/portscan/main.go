package main

import (
	"fmt"

	"facemasq/lib/scans"
	scanPorts "facemasq/lib/scans/port"
)

func main() {
	inputs := scans.AddressesToPortScan{
		scans.AddressToPortScan{
			AddressID: 12,
			IPv4:      "192.168.0.24",
		},
		scans.AddressToPortScan{
			AddressID: 7,
			IPv4:      "192.168.0.20",
		},
		scans.AddressToPortScan{
			AddressID: 62,
			IPv4:      "192.168.0.41",
		},
	}

	for i := range inputs {

		scannedPorts := scanPorts.ScanAddressAsync(inputs[i].IPv4)
		fmt.Printf("%s Open TCP ports: ", inputs[i].IPv4)
		for j := range scannedPorts.Ports {
			if scannedPorts.Ports[j].Protocol == "tcp" {
				fmt.Printf("%d, ", scannedPorts.Ports[j].Number)
			}
		}
		// fmt.Println()
		// fmt.Printf("%s Open UDP ports:\n", inputs[i].IPv4)
		// for j := range scannedPorts.Ports {
		// 	if scannedPorts.Ports[j].Protocol == "udp" {
		// 		fmt.Printf("%d,", scannedPorts.Ports[j].Number)
		// 	}
		// }
		fmt.Println()
	}
}

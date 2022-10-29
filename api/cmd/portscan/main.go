package main

import (
	"fmt"
	"time"
)

func main() {
	// inputs := scans.AddressesToPortScan{
	// 	scans.AddressToPortScan{
	// 		AddressID: 12,
	// 		IPv4:      "192.168.3.231",
	// 	},
	// }

	// for i := range inputs {

	// 	scannedPorts := scanPorts.ScanAddress(inputs[i].IPv4)
	// 	fmt.Printf("%s Open TCP ports: ", inputs[i].IPv4)
	// 	for j := range scannedPorts.Ports {
	// 		if scannedPorts.Ports[j].Protocol == "tcp" {
	// 			fmt.Printf("%d, ", scannedPorts.Ports[j].Number)
	// 		}
	// 	}
	// 	// fmt.Println()
	// 	// fmt.Printf("%s Open UDP ports:\n", inputs[i].IPv4)
	// 	// for j := range scannedPorts.Ports {
	// 	// 	if scannedPorts.Ports[j].Protocol == "udp" {
	// 	// 		fmt.Printf("%d,", scannedPorts.Ports[j].Number)
	// 	// 	}
	// 	// }
	// 	fmt.Println()
	// }
	var portList []int64
	for i := int64(1); i <= 1024; i++ {
		portList = append(portList, i)
	}

	for a := 0; a < 3; a++ {
		j := 0
		for i := range portList {
			if i != 0 && i%(len(portList)/10) == 0 {
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("%d", j)
				j++

			}
		}
		fmt.Println()

	}
}

package port

import (
	"context"
	"facemasq/lib/logging"
	"facemasq/lib/network"
	"facemasq/lib/scans"
	"facemasq/models"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"testing"
	"time"

	"github.com/liamg/furious/scan"
	"github.com/volatiletech/null"
)

func TestPortScan(t *testing.T) {

	records, err := getRecords(1)
	if err != nil {
		t.Fatal(err)
	}

	maxscans := 5
	var ports []models.Port
	t.Logf("%d addresses available to scan", len(records))
	for i := range records {
		t.Logf("address #%d = %s ", i, records[i].IPv4)
		if i >= maxscans {
			t.Logf("Iteration %d meets or exceeds maxscans %d ", i, maxscans)
			break
		}
		result := ScanAddress(records[i].IPv4)
		for _, result := range result.Ports {
			ports = append(ports, models.Port{AddressID: int64(records[i].AddressID), ScanID: 1, Port: result.Number, Protocol: result.Protocol})
		}
	}
	for p := range ports {
		t.Logf("%+v", ports[p])

	}
}

func getRecords(scanID int64) (records scans.DeviceRecords, err error) {
	lastSeen := time.Now().Format("2006-01-02 15:04:05")
	// Get details of local interfaces
	logging.Println(1, "Processing local interfaces")
	records, err = getLocalIFaces(scanID, lastSeen)
	if err != nil {
		err = fmt.Errorf("could not get local interfaces: %v", err)
		scanID = 0
		return
	}

	// Scan the $target network
	logging.Println(1, "Scanning network")
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancel()
	}()
	timeOutDuration := time.Duration(timeOut) * time.Millisecond
	targetIterator := scan.NewTargetIterator(network.Target)
	scanner := scan.NewDeviceScanner(targetIterator, timeOutDuration)

	err = scanner.Start()
	if err != nil {
		err = fmt.Errorf("could not initialise scanner: %v", err)
		return
	}
	results, err := scanner.Scan(ctx, scan.DefaultPorts)
	if err != nil {
		err = fmt.Errorf("could not complete scan: %v", err)
		return
	}
	for _, result := range results {
		if result.MAC != "" { // && result.IsHostUp() {
			record := scans.DeviceRecord{
				ScanID:   scanID,
				MAC:      strings.ToUpper(result.MAC),
				IPv4:     result.Host.String(),
				LastSeen: lastSeen,
				Hostname: result.Name,
				Notes:    result.Manufacturer,
			}

			logging.Printf(3, "Found %s (%s) via NET", record.IPv4, record.MAC)

			records = append(records, record)
		}
	}
	return
}

func getLocalIFaces(scanID int64, lastSeen string) (records scans.DeviceRecords, err error) {
	var record scans.DeviceRecord
	var ipv6 null.String
	var ipv4 string

	netFaces, err := net.Interfaces()
	if err != nil {
		logging.Errorf("ProcessLocal: %+v", err.Error())
		return
	}
	for _, netFace := range netFaces {
		if !strings.Contains(netFace.Name, "veth") && !strings.Contains(netFace.Name, "lo") && !strings.Contains(netFace.Name, "br-") && !strings.Contains(netFace.Name, "docker0") {
			addresses, err := netFace.Addrs()
			if err != nil {
				logging.Errorf("ProcessLocal: %+v", err.Error())
				continue
			}
			if len(addresses) > 0 {
				ipv6 = null.NewString("", false)
				ipv4 = ""
				for _, address := range addresses {

					IP := strings.Split(address.String(), "/")[0]
					if strings.Contains(IP, ":") {
						ipv6 = null.StringFrom(IP)
					} else {
						ipv4 = IP
					}

				}
				record = scans.DeviceRecord{
					ScanID:   scanID,
					MAC:      strings.ToUpper(netFace.HardwareAddr.String()),
					IPv4:     ipv4,
					IPv6:     ipv6,
					LastSeen: lastSeen,
					Hostname: "",
					Notes:    "",
				}
				records = append(records, record)
			}
		}
	}
	return
}

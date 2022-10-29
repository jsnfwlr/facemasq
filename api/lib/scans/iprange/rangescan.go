package iprange

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/liamg/furious/scan"
	"github.com/volatiletech/null"

	"facemasq/lib/db"
	"facemasq/lib/logging"
	"facemasq/lib/macvendor"
	"facemasq/lib/network"
	"facemasq/lib/scans"
	portscan "facemasq/lib/scans/port"
	"facemasq/models"
)

var (
	Frequency time.Duration
	timeOut   int
)

var UseScanLog = false

func init() {
	var err error
	Frequency, err = time.ParseDuration(os.Getenv("NETSCAN_FREQUENCY"))
	if err != nil {
		Frequency = time.Duration(60) * time.Second
	}
	timeOut, err = strconv.Atoi(os.Getenv("NETSCAN_TIMEOUT"))
	if err != nil {
		timeOut = 2000
	}
}

func Schedule() {
	sched := gocron.NewScheduler(time.UTC)
	sched.Every(Frequency).Do(func() {
		scanID, err := ScanAndStore()
		if err != nil {
			logging.Printf(0, "error: %v", err)
		}
		err = portscan.ScanAysnc(scanID, false)
		if err != nil {
			logging.Printf(0, "%v", err)
		}
	})
	sched.StartAsync()
}

func Scan(scanID int64) (records scans.DeviceRecords, err error) {
	macvendor.TooManyRequests = false
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

			logging.Printf(3, "Found %s (%s) via NET\n", record.IPv4, record.MAC)

			records = append(records, record)
		}
	}
	return
}

func Store(records scans.DeviceRecords, scanID int64) (err error) {
	err = records.Store()
	if err != nil {
		logging.Printf(0, "error recording netscan results: %v", err)
		err = nil // a single error here shouldn't stop the whole process
	}

	// Mark all devices as inactive, then set all the inactive devices from the recent scan to active
	sqlQ := `UPDATE devices SET is_online = false WHERE is_online = true;`
	_, err = db.Conn.Exec(sqlQ)
	if err != nil {
		logging.Printf(0, "could not update device statuses: %v\n", err)
		// return
	}

	sqlQ = `UPDATE interfaces SET is_online = false WHERE is_online = true;`
	_, err = db.Conn.Exec(sqlQ)
	if err != nil {
		logging.Printf(0, "could not update Interfaces statuses: %v\n", err)
		// return
	}

	sqlQ = `UPDATE devices SET is_online = true WHERE is_online = false AND id IN (SELECT interfaces.device_id FROM interfaces JOIN addresses ON addresses.interface_id = interfaces.id JOIN histories ON histories.address_id = addresses.id AND histories.scan_id = ?);`
	_, err = db.Conn.Exec(sqlQ, scanID)
	if err != nil {
		logging.Printf(0, "could not update device statuses: %v\n", err)
		// return
	}

	sqlQ = `UPDATE interfaces SET is_online = true WHERE is_online = false AND ID IN (SELECT addresses.interface_id FROM addresses JOIN histories ON histories.address_id = addresses.id AND histories.scan_id = ?);`
	_, err = db.Conn.Exec(sqlQ, scanID)
	if err != nil {
		logging.Printf(0, "could not updatie Interfaces statuses: %v\n", err)
		// return
	}
	return
}

func ScanAndStore() (scanID int64, err error) {
	scanRecord := models.Scan{
		Time: time.Now(),
	}
	_, err = db.Conn.NewInsert().Model(&scanRecord).Exec(db.Context)
	if err != nil {
		return
	}
	scanID = scanRecord.ID

	var records scans.DeviceRecords
	records, err = Scan(scanID)
	if err != nil {
		return
	}
	err = Store(records, scanID)

	return
}

func getLocalIFaces(scanID int64, lastSeen string) (records scans.DeviceRecords, err error) {
	var record scans.DeviceRecord
	var ipv6 null.String
	var ipv4 string

	netFaces, err := net.Interfaces()
	if err != nil {
		logging.Printf(0, "ProcessLocal: %+v\n", err.Error())
		return
	}
	for _, netFace := range netFaces {
		if !strings.Contains(netFace.Name, "veth") && !strings.Contains(netFace.Name, "lo") && !strings.Contains(netFace.Name, "br-") && !strings.Contains(netFace.Name, "docker0") {
			addresses, err := netFace.Addrs()
			if err != nil {
				logging.Printf(0, "ProcessLocal: %+v\n", err.Error())
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

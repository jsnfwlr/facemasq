package netscan

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/liamg/furious/scan"
	"github.com/volatiletech/null"

	"github.com/jsnfwlr/facemasq/api/lib/db"
	"github.com/jsnfwlr/facemasq/api/models/Addresses"
	"github.com/jsnfwlr/facemasq/api/models/Netfaces"
)

var target string
var frequency time.Duration
var timeOut int

const UseScanLog = false

func init() {
	var err error
	target = os.Getenv("TARGET")
	frequency, err = time.ParseDuration(os.Getenv("FREQUENCY"))
	if err != nil {
		frequency = time.Duration(1) * time.Minute
	}
	timeOut, err = strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		timeOut = 2000
	}
}

func Schedule() {

	sched := gocron.NewScheduler(time.UTC)
	ShowNetworkSummary()
	sched.Every(frequency).Do(func() {
		err := ScanAndStore()
		if err != nil {
			fmt.Printf("error: %v", err)
		}
	})
	sched.StartAsync()
}

func ShowNetworkSummary() (err error) {
	var (
		addresses                []string
		network, broadcast, mask string
	)
	addresses, network, broadcast, mask, err = getIPRange(target)
	if err != nil {
		return
	}
	log.Printf("Network: %s, Broadcast: %s, Mask: %s, Addresses: %d\n", network, broadcast, mask, len(addresses))
	return
}

func getLocalIFaces(scanID int64, lastSeen string) (records []Result, err error) {
	var record Result
	var ipv6 null.String
	var ipv4 string

	netFaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("ProcessLocal: %+v\n", err.Error())
		return
	}
	for _, netFace := range netFaces {
		if !strings.Contains(netFace.Name, "veth") && !strings.Contains(netFace.Name, "lo") && !strings.Contains(netFace.Name, "br-") && !strings.Contains(netFace.Name, "docker0") {
			addresses, err := netFace.Addrs()
			if err != nil {
				fmt.Printf("ProcessLocal: %+v\n", err.Error())
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
				record = Result{
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

func ScanAndStore() (err error) {

	var sqlres sql.Result
	var records []Result

	// var histories []models.History

	lastSeen := time.Now().Format("2006-01-02 15:04:05")

	sql := `INSERT INTO Scans (Time) VALUES (?);`
	sqlres, err = db.Conn.Exec(sql, lastSeen)
	if err != nil {
		return
	}
	scanID, _ := sqlres.LastInsertId()

	// Get details of local interfaces
	log.Println("Processing local interfaces")
	records, err = getLocalIFaces(scanID, lastSeen)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, record := range records {
		err = record.Store()
		if err != nil {
			return
		}
	}

	// Scan the $target network
	log.Println("Scanning network")

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancel()
	}()
	timeOutDuration := time.Duration(timeOut) * time.Millisecond
	targetIterator := scan.NewTargetIterator(target)
	scanner := scan.NewDeviceScanner(targetIterator, timeOutDuration)

	err = scanner.Start()
	if err != nil {
		return
	}
	results, err := scanner.Scan(ctx, scan.DefaultPorts)
	if err != nil {
		return
	}

	for _, result := range results {
		if result.MAC != "" { // && result.IsHostUp() {
			record := Result{
				ScanID:   scanID,
				MAC:      strings.ToUpper(result.MAC),
				IPv4:     result.Host.String(),
				LastSeen: lastSeen,
				Hostname: result.Name,
				Notes:    result.Manufacturer,
			}
			err = record.Store()
			if err != nil {
				return
			}
		}
	}

	// Mark all devices as inactive, then set all the inactive devices from the recent scan to active
	sql = `UPDATE Devices SET IsOnline = false WHERE IsOnline = true;`
	_, err = db.Conn.Exec(sql)
	if err != nil {
		fmt.Println("Error updating device statuses")
		return
	}

	sql = `UPDATE Interfaces SET IsOnline = false WHERE IsOnline = true;`
	_, err = db.Conn.Exec(sql)
	if err != nil {
		fmt.Println("Error updating Interfaces statuses")
		return
	}

	sql = `UPDATE Devices SET IsOnline = true WHERE IsOnline = false AND ID IN (SELECT Interfaces.DeviceID FROM Interfaces JOIN Addresses ON Addresses.InterfaceID = Interfaces.ID JOIN History ON History.AddressID = Addresses.ID AND History.ScanID = ?);`
	_, err = db.Conn.Exec(sql, scanID)
	if err != nil {
		fmt.Println("Error updating device statuses")
		return
	}

	sql = `UPDATE Interfaces SET IsOnline = true WHERE IsOnline = false AND ID IN (SELECT Addresses.InterfaceID FROM Addresses JOIN History ON History.AddressID = Addresses.ID AND History.ScanID = ?);`
	_, err = db.Conn.Exec(sql, scanID)
	if err != nil {
		fmt.Println("Error updating Interfaces statuses")
		return
	}
	return
}

func (record Result) Store() (err error) {
	var netface Netfaces.Model
	var address Addresses.Model
	var sqlres sql.Result

	sql := `SELECT * FROM Interfaces WHERE MAC = ?;`
	err = db.Conn.Get(&netface, sql, record.MAC)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			fmt.Printf("%v\n", err)
			return
		}
		sql = `INSERT INTO Devices (FirstSeen, Notes, MachineName, Label, IsOnline, StatusID) VALUES (?, ?, "", "Unknown Device", true, 1)`
		sqlres, err = db.Conn.Exec(sql, record.LastSeen, record.Notes)
		if err != nil {
			fmt.Println("Error creating device")
			return
		}
		lastDeviceID, _ := sqlres.LastInsertId()

		// sql = `SELECT ID FROM Device WHERE Label = ?`
		// err = db.Conn.Get(&device, sql, record.MAC)
		// if err != nil {
		// 	return
		// }

		sql = `INSERT INTO Interfaces (MAC, DeviceID, Notes, LastSeen, IsOnline, StatusID) VALUES (?, ?, ?, ?, true, 1);`
		sqlres, err = db.Conn.Exec(sql, record.MAC, lastDeviceID, record.Notes, record.LastSeen)
		if err != nil {
			fmt.Println("Error creating interface")
			return
		}
		lastInterfaceID, _ := sqlres.LastInsertId()

		sql = `INSERT INTO Addresses (IPv4, IPv6, InterfaceID, LastSeen, Notes) VALUES (?, ?, ?, ?, ?);`
		sqlres, err = db.Conn.Exec(sql, record.IPv4, record.IPv6, lastInterfaceID, record.LastSeen, record.Notes)
		if err != nil {
			fmt.Println("Error creating address")
			return
		}
		lastAddressID, _ := sqlres.LastInsertId()

		if record.Hostname != "" {
			sql = `INSERT INTO Hostnames (Hostname, AddressID, Notes) VALUES (?,?, ?);`
			_, err = db.Conn.Exec(sql, record.Hostname, lastAddressID, record.Notes)
			if err != nil {
				fmt.Println("Error creating hostname")
				return
			}
		}

		sql = `INSERT INTO History (AddressID, ScanID) VALUES (?,?);`
		_, err = db.Conn.Exec(sql, lastAddressID, record.ScanID)
		if err != nil {
			fmt.Println("Error creating history")
			return
		}
		return
	}

	sql = `SELECT * FROM Addresses WHERE IPv4 = ? AND InterfaceID = ?`
	err = db.Conn.Get(&address, sql, record.IPv4, netface.ID)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			fmt.Printf("%v\n", err)
			return
		}

		sql = `INSERT INTO Addresses (IPv4, IPv6, InterfaceID, LastSeen, Notes) VALUES (?, ?, ?, ?, ?);`
		sqlres, err = db.Conn.Exec(sql, record.IPv4, record.IPv6, netface.ID, record.LastSeen, record.Notes)
		if err != nil {
			fmt.Println("Error adding address")
			return
		}
		lastAddressID, _ := sqlres.LastInsertId()

		if record.Hostname != "" {
			sql = `INSERT INTO Hostnames (Hostname, AddressID, Notes) VALUES (?, ?, ?);`
			_, err = db.Conn.Exec(sql, record.Hostname, lastAddressID, record.Notes)
			if err != nil {
				fmt.Println("Error adding hostname")
				return
			}
		}

		sql = `INSERT INTO History (AddressID, ScanID) VALUES (?, ?);`
		_, err = db.Conn.Exec(sql, lastAddressID, record.ScanID)
		if err != nil {
			fmt.Println("Error adding history")
			return
		}

		sql = `UPDATE Addresses SET Label = NULL WHERE Label = ?;`
		_, err = db.Conn.Exec(sql, record.MAC)
		if err != nil {
			return
		}
		return
	}

	sql = `UPDATE Interfaces SET LastSeen = ? WHERE ID = ?;`
	_, err = db.Conn.Exec(sql, record.LastSeen, netface.ID)
	if err != nil {
		fmt.Println("Error recording interfaces lastseen")
		return
	}

	sql = `UPDATE Addresses SET LastSeen = ? WHERE ID = ?;`
	_, err = db.Conn.Exec(sql, record.LastSeen, address.ID)
	if err != nil {
		fmt.Println("Error recording address lastseen")
		return
	}

	sql = `INSERT INTO History (AddressID, ScanID) VALUES (?,?);`
	_, err = db.Conn.Exec(sql, address.ID, record.ScanID)
	if err != nil {
		fmt.Println("Error recording new history")
		return
	}
	return
}

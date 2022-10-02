package scanresults

import (
	"fmt"
	"time"

	"facemasq/lib/db"
	"facemasq/lib/devices"
	"facemasq/lib/logging"
	"facemasq/lib/macvendor"
	"facemasq/models"

	"github.com/uptrace/bun"
	"github.com/volatiletech/null"
)

func (records Records) GroupParams() (ipv4, mac []string) {
	for i := range records {
		ipv4 = append(ipv4, records[i].IPv4)
		mac = append(mac, records[i].MAC)
	}
	return
}

func getDevices() (allDevices []models.Device, err error) {
	queries := devices.DeviceQueries{
		Devices:   `SELECT * FROM devices;`,
		Netfaces:  `SELECT * FROM interfaces ORDER BY is_primary DESC, is_virtual ASC;`,
		Addresses: `SELECT * FROM addresses ORDER BY interface_id ASC, is_primary DESC, last_seen DESC, is_virtual ASC;`,
		Hostnames: `SELECT * FROM hostnames;`,
	}
	allDevices, err = devices.GetDevices(queries, "", false)
	return
}

func (records Records) filterNew(allDevices []models.Device) (newAddressRecords, newDeviceRecords []Record) {
	for r := range records {
		newInterface := true
		newAddress := true
		for d := range allDevices {
			for i := range allDevices[d].Interfaces {
				if allDevices[d].Interfaces[i].MAC == records[r].MAC {
					newInterface = false

					for a := range allDevices[d].Interfaces[i].Addresses {
						if allDevices[d].Interfaces[i].Addresses[a].IPv4.String == records[r].IPv4 {
							newAddress = false
						}
					}
					if newAddress {
						records[r].InterfaceID = allDevices[d].Interfaces[i].ID
						newAddressRecords = append(newAddressRecords, records[r])
					}
				}
			}
		}
		if newInterface {
			newDeviceRecords = append(newDeviceRecords, records[r])
		}
	}
	return
}

func (records Records) Store() (err error) {
	var allDevices []models.Device

	allDevices, err = getDevices()
	if err != nil {
		return
	}

	newAddressRecords, newDeviceRecords := records.filterNew(allDevices)

	if len(newAddressRecords) > 0 {
		for a := range newAddressRecords {
			err = newAddressRecords[a].CreateAddress()
			if err != nil {
				return
			}
		}
	}
	if len(newDeviceRecords) > 0 {
		for d := range newDeviceRecords {
			err = newDeviceRecords[d].CreateDevice()
			if err != nil {
				return
			}
		}
	}

	ipv4, mac := records.GroupParams()

	logging.Printf(1, "Updating %d Interfaces\n", len(mac))

	_, err = db.Conn.NewUpdate().Model((*models.Interface)(nil)).Set("last_seen = ?", records[0].LastSeen).Where("mac IN (?)", bun.In(mac)).Exec(db.Context)
	if err != nil {
		err = fmt.Errorf("could not bulk-update Inteface last_seen: %v", err)
		logging.Printf(0, "Err with bulk update of interfaces %v\n", err)
		return
	}

	logging.Printf(1, "Updating %d Addresses\n", len(ipv4))
	_, err = db.Conn.NewUpdate().Table("addresses").Set("last_seen = ?", records[0].LastSeen).Where("ipv4 IN (?)", bun.In(ipv4)).Where("interface_id IN (SELECT id FROM interfaces WHERE mac IN (?))", bun.In(mac)).Exec(db.Context)
	if err != nil {
		err = fmt.Errorf("could not bulk-update Address last_seen: %v", err)
		return
	}

	logging.Printf(1, "Updating %d History\n", len(ipv4))
	var history []models.History
	err = db.Conn.NewRaw(`SELECT id AS address_id, ? AS scan_id FROM addresses WHERE ipv4 IN (?) AND interface_id IN (SELECT id FROM interfaces WHERE mac IN (?));`, records[0].ScanID, bun.In(ipv4), bun.In(mac)).Scan(db.Context, &history)
	if err != nil {
		err = fmt.Errorf("could not generate bulk-insert history: %v", err)
		logging.Printf(0, "could not generate bulk-insert history %v\n", err)
		return
	}
	_, err = db.Conn.NewInsert().Model(&history).Exec(db.Context)
	if err != nil {
		err = fmt.Errorf("could not bulk-insert history: %v", err)
		logging.Printf(0, "could not bulk-insert history: %v\n", err)
		return
	}
	return
}

func (record *Record) CreateDevice() (err error) {
	device := models.Device{
		MachineName: "",
		Label:       null.String{String: "Unknown Device", Valid: true},
		IsOnline:    null.BoolFrom(true),
		Notes:       null.String{String: record.Notes, Valid: true},
		StatusID:    1,
		FirstSeen:   time.Now(),
	}

	vendor, err := macvendor.Lookup(record.MAC)
	if err != nil {
		logging.Printf(0, "could not lookup vendor for MAC Address (%s): %v\n", record.MAC, err)
		err = nil
	}
	if vendor != "" {
		device.Brand = null.StringFrom(vendor)
	}
	_, err = db.Conn.NewInsert().Model(&device).Exec(db.Context)
	if err != nil {
		return
	}
	record.DeviceID = device.ID
	err = record.CreateInterface()
	return
}

func (record *Record) CreateInterface() (err error) {
	netface := models.Interface{
		MAC:      record.MAC,
		StatusID: 1,
		IsOnline: true,
		Notes:    null.String{String: record.Notes, Valid: true},
		DeviceID: record.DeviceID,
	}
	_, err = db.Conn.NewInsert().Model(&netface).Exec(db.Context)
	if err != nil {
		return
	}
	record.InterfaceID = netface.ID
	err = record.CreateAddress()
	return
}

func (record *Record) CreateAddress() (err error) {
	address := models.Address{
		InterfaceID: record.InterfaceID,
		Notes:       null.String{String: record.Notes, Valid: true},
		IPv4:        null.String{String: record.IPv4, Valid: true},
	}
	_, err = db.Conn.NewInsert().Model(&address).Exec(db.Context)
	if err != nil {
		return
	}
	record.AddressID = address.ID
	if record.Hostname != "" {
		err = record.CreateHostname()
	}
	return
}

func (record *Record) CreateHostname() (err error) {
	hostname := models.Hostname{
		AddressID: record.AddressID,
		Notes:     null.String{String: record.Notes, Valid: true},
		Hostname:  record.Hostname,
	}
	_, err = db.Conn.NewInsert().Model(&hostname).Exec(db.Context)
	if err != nil {
		return
	}
	return
}

func PrepareForTest() (err error) {
	scan := models.Scan{
		Time: time.Now(),
	}
	_, err = db.Conn.NewInsert().Model(&scan).Exec(db.Context)
	if err != nil {
		return
	}
	records := Records{
		{
			ScanID: scan.ID,
			IPv4:   "192.168.0.1",
			MAC:    "00:00:00:00:00:00:01:00",
		},
		{
			ScanID: scan.ID,
			IPv4:   "192.168.0.2",
			MAC:    "00:00:00:00:00:00:01:01",
		},
		{
			ScanID:   scan.ID,
			Hostname: "TestDevice",
			IPv4:     "192.168.0.3",
			MAC:      "00:00:00:00:00:00:01:02",
		},
	}

	err = records.Store()

	return

}

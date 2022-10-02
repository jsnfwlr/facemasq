package db

import (
	"facemasq/lib/files"
	"facemasq/lib/logging"
	"facemasq/models"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/volatiletech/null"
)

func ConnectToTest() (err error) {
	var dataPath string

	switch strings.ToLower(DBEngine) {
	case "sqlite":
		dataPath, err = files.GetDir("data")
		if err != nil {
			logging.Panic(err)
		}

		DBConnString = fmt.Sprintf("file:%[2]s%[1]c%[3]s", os.PathSeparator, dataPath, "network.sqlite")
	case "postgres":
		DBConnString = "pg://phalacee:h34rts@(192.168.0.41)/test_facemasq"
	case "mysql":
		DBConnString = "phalacee:h34rts@(192.168.0.41)/test_facemasq"
	}
	err = Connect()
	if err != nil {
		return
	}
	return preloadTestData()
}

func preloadTestData() (err error) {

	lastSeen := time.Now()

	devices := []models.Device{
		{
			Label:     null.String{String: "Testing 1", Valid: true},
			Notes:     null.String{String: "Testing", Valid: true},
			FirstSeen: time.Now(),
		},
		{
			Label:     null.String{String: "Testing 2", Valid: true},
			Notes:     null.String{String: "Testing", Valid: true},
			FirstSeen: time.Now(),
		},
		{
			Label:     null.String{String: "Testing 3", Valid: true},
			Notes:     null.String{String: "Testing", Valid: true},
			FirstSeen: time.Now(),
		},
	}
	_, err = Conn.NewInsert().Model(&devices).Exec(Context)
	if err != nil {
		return
	}

	netfaces := []models.Interface{
		{
			MAC:             "00:00:00:00:00:00",
			IsPrimary:       true,
			IsVirtual:       false,
			InterfaceTypeID: 1,
			LastSeen:        lastSeen,
			VlanID:          1,
			DeviceID:        devices[0].ID,
			Label:           null.String{String: "Test", Valid: true},
			Notes:           null.String{String: "Testing", Valid: true},
		},
		{
			MAC:             "00:00:00:00:00:01",
			IsPrimary:       true,
			IsVirtual:       false,
			InterfaceTypeID: 1,
			LastSeen:        lastSeen,
			VlanID:          1,
			DeviceID:        devices[1].ID,
			Label:           null.String{String: "Test", Valid: true},
			Notes:           null.String{String: "Testing", Valid: true},
		},
		{
			MAC:             "00:00:00:00:01:00",
			IsPrimary:       true,
			IsVirtual:       false,
			InterfaceTypeID: 1,
			LastSeen:        lastSeen,
			VlanID:          1,
			DeviceID:        devices[2].ID,
			Label:           null.String{String: "Test", Valid: true},
			Notes:           null.String{String: "Testing", Valid: true},
		},
		{
			MAC:             "00:00:00:00:01:01",
			IsPrimary:       true,
			IsVirtual:       false,
			InterfaceTypeID: 1,
			LastSeen:        lastSeen,
			VlanID:          1,
			DeviceID:        devices[2].ID,
			Label:           null.String{String: "Test", Valid: true},
			Notes:           null.String{String: "Testing", Valid: true},
		},
	}
	_, err = Conn.NewInsert().Model(&netfaces).Exec(Context)
	if err != nil {
		return
	}

	addresses := []models.Address{
		{
			IPv4:        null.String{String: "192.168.0.1", Valid: true},
			IPv6:        null.String{String: ""},
			IsPrimary:   null.BoolFrom(true),
			IsVirtual:   null.BoolFrom(false),
			IsReserved:  null.BoolFrom(false),
			LastSeen:    time.Now(),
			Label:       null.String{String: "Test", Valid: true},
			Notes:       null.String{String: "This should be deleted", Valid: true},
			InterfaceID: netfaces[0].ID,
		},
		{
			IPv4:        null.String{String: "192.168.0.2", Valid: true},
			IPv6:        null.String{String: ""},
			IsPrimary:   null.BoolFrom(true),
			IsVirtual:   null.BoolFrom(false),
			IsReserved:  null.BoolFrom(false),
			LastSeen:    time.Now(),
			Label:       null.String{String: "Test", Valid: true},
			Notes:       null.String{String: "This should be deleted", Valid: true},
			InterfaceID: netfaces[1].ID,
		},
		{
			IPv4:        null.String{String: "192.168.1.1", Valid: true},
			IPv6:        null.String{String: ""},
			IsPrimary:   null.BoolFrom(true),
			IsVirtual:   null.BoolFrom(false),
			IsReserved:  null.BoolFrom(false),
			LastSeen:    time.Now(),
			Label:       null.String{String: "Test", Valid: true},
			Notes:       null.String{String: "This should be deleted", Valid: true},
			InterfaceID: netfaces[2].ID,
		},
		{
			IPv4:        null.String{String: "192.168.1.2", Valid: true},
			IPv6:        null.String{String: ""},
			IsPrimary:   null.BoolFrom(true),
			IsVirtual:   null.BoolFrom(false),
			IsReserved:  null.BoolFrom(false),
			LastSeen:    time.Now(),
			Label:       null.String{String: "Test", Valid: true},
			Notes:       null.String{String: "This should be deleted", Valid: true},
			InterfaceID: netfaces[3].ID,
		},
		{
			IPv4:        null.String{String: "192.168.1.3", Valid: true},
			IPv6:        null.String{String: ""},
			IsPrimary:   null.BoolFrom(true),
			IsVirtual:   null.BoolFrom(false),
			IsReserved:  null.BoolFrom(false),
			LastSeen:    time.Now(),
			Label:       null.String{String: "Test", Valid: true},
			Notes:       null.String{String: "This should be deleted", Valid: true},
			InterfaceID: netfaces[3].ID,
		},
	}
	_, err = Conn.NewInsert().Model(&addresses).Exec(Context)
	if err != nil {
		return
	}

	hostnames := []models.Hostname{
		{
			Hostname:  "Host0",
			AddressID: addresses[0].ID,
			IsDNS:     true,
			IsSelfSet: false,
		},
		{
			Hostname:  "Host1",
			AddressID: addresses[1].ID,
			IsDNS:     true,
			IsSelfSet: false,
		},
	}
	_, err = Conn.NewInsert().Model(&hostnames).Exec(Context)
	if err != nil {
		return
	}

	return
}

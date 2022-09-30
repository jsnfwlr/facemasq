package scanresults

import (
	"fmt"
	"os"
	"testing"
	"time"

	"facemasq/lib/db"
	"facemasq/lib/files"
	"facemasq/lib/logging"
	"facemasq/models"

	"github.com/volatiletech/null"
)

func TestBulkStore(test *testing.T) {
	lastSeen := time.Now()

	db.DBEngine = "sqlite"
	dataPath, _ := files.GetDir("data")
	db.DBConnString = fmt.Sprintf("file:%[2]s%[1]c%[3]s", os.PathSeparator, dataPath, "test2.sqlite")

	err := db.Connect()
	if err != nil {
		test.Error(err)
	}
	device1 := models.Device{
		Label:     null.String{String: "Testing 1", Valid: true},
		Notes:     null.String{String: "Testing", Valid: true},
		FirstSeen: time.Now(),
	}
	device2 := models.Device{
		Label:     null.String{String: "Testing 2", Valid: true},
		Notes:     null.String{String: "Testing", Valid: true},
		FirstSeen: time.Now(),
	}
	_, err = db.Conn.NewInsert().Model(&device1).Exec(db.Context)
	if err != nil {
		test.Error(err)
	}
	_, err = db.Conn.NewInsert().Model(&device2).Exec(db.Context)
	if err != nil {
		test.Error(err)
	}

	netface1 := models.Interface{
		MAC:             "00:00:00:00:00:00:00:00",
		IsPrimary:       true,
		IsVirtual:       false,
		InterfaceTypeID: 1,
		LastSeen:        lastSeen,
		VlanID:          1,
		DeviceID:        1,
		Label:           null.String{String: "Test", Valid: true},
		Notes:           null.String{String: "Testing", Valid: true},
	}
	netface2 := models.Interface{
		MAC:             "00:00:00:00:00:00:00:01",
		IsPrimary:       true,
		IsVirtual:       false,
		InterfaceTypeID: 1,
		LastSeen:        lastSeen,
		VlanID:          1,
		DeviceID:        2,
		Label:           null.String{String: "Test", Valid: true},
		Notes:           null.String{String: "Testing", Valid: true},
	}
	_, err = db.Conn.NewInsert().Model(&netface1).Exec(db.Context)
	if err != nil {
		test.Error(err)
	}
	_, err = db.Conn.NewInsert().Model(&netface2).Exec(db.Context)
	if err != nil {
		test.Error(err)
	}

	address1 := models.Address{
		IPv4:        null.String{String: "192.168.5.1", Valid: true},
		IPv6:        null.String{String: ""},
		IsPrimary:   null.BoolFrom(true),
		IsVirtual:   null.BoolFrom(false),
		IsReserved:  null.BoolFrom(false),
		LastSeen:    time.Now(),
		Label:       null.String{String: "Test", Valid: true},
		Notes:       null.String{String: "This should be deleted", Valid: true},
		InterfaceID: 1,
	}
	address2 := models.Address{
		IPv4:        null.String{String: "192.168.4.1", Valid: true},
		IPv6:        null.String{String: ""},
		IsPrimary:   null.BoolFrom(true),
		IsVirtual:   null.BoolFrom(false),
		IsReserved:  null.BoolFrom(false),
		LastSeen:    time.Now(),
		Label:       null.String{String: "Test", Valid: true},
		Notes:       null.String{String: "This should be deleted", Valid: true},
		InterfaceID: 2,
	}
	_, err = db.Conn.NewInsert().Model(&address1).Exec(db.Context)
	if err != nil {
		test.Error(err)
	}
	_, err = db.Conn.NewInsert().Model(&address2).Exec(db.Context)
	if err != nil {
		test.Error(err)
	}

	scan := models.Scan{
		Time: lastSeen,
	}
	_, err = db.Conn.NewInsert().Model(&scan).Exec(db.Context)
	if err != nil {
		test.Error(err)
	}

	results := Records{
		Record{
			ScanID: scan.ID,
			IPv4:   address1.IPv4.String,
			MAC:    netface1.MAC,
		},
		Record{
			ScanID: scan.ID,
			IPv4:   address2.IPv4.String,
			MAC:    netface2.MAC,
		},
	}
	logging.Verbosity = 2
	err = results.Store()
	if err != nil {
		test.Error(err)
	}
}

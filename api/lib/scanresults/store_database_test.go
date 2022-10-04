//go:build database || full

package scanresults

import (
	"testing"
	"time"

	"facemasq/lib/db"
	"facemasq/lib/logging"
	"facemasq/models"
)

func TestBulkStore(t *testing.T) {
	lastSeen := time.Now()

	container, err := db.ConnectToTest()
	if err != nil {
		t.Fatal(err)
	}
	defer container.Close()

	scan := models.Scan{
		Time: lastSeen,
	}
	_, err = db.Conn.NewInsert().Model(&scan).Exec(db.Context)
	if err != nil {
		t.Fatal(err)
	}

	results := Records{
		// {
		// 	ScanID: scan.ID,
		// 	IPv4:   address1.IPv4.String,
		// 	MAC:    netface1.MAC,
		// },
		// {
		// 	ScanID: scan.ID,
		// 	IPv4:   address2.IPv4.String,
		// 	MAC:    netface2.MAC,
		// },
		{
			ScanID:   scan.ID,
			Hostname: "TestDevice",
			IPv4:     "192.168.3.1",
			MAC:      "00:00:00:00:00:00:00:02",
		},
	}
	logging.Verbosity = 2
	err = results.Store()
	if err != nil {
		t.Error(err)
	}
	err = container.Close()
	if err != nil {
		t.Errorf("Error closing container: %v", err)
	}
}

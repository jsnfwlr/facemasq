//go:build full || database

package iprange

import (
	"facemasq/lib/db"
	portscan "facemasq/lib/scans/port"
	"facemasq/models"
	"testing"
)

func TestScanAndStore(t *testing.T) {
	t.Log("ScanAndStore")
	// for e := range db.DBEngines {
	// 	t.Log(db.DBEngines[e])
	db.DBEngine = "mysql" // db.DBEngines[e]

	cntnr, err := db.ConnectToTest(false)
	if err != nil {
		t.Fatal(err)
	}
	db.RunningContainer = cntnr

	scanID, err := ScanAndStore()
	if err != nil {
		t.Error(err)
	}

	err = portscan.Scan(scanID, true)
	if err != nil {
		t.Error(err)
	}

	var ports []models.Port
	err = db.Conn.NewSelect().Model(&ports).Scan(db.Context)
	if err != nil {
		t.Error(err)
	}
	for i := range ports {
		t.Logf("port: %+v", ports[i])
	}

	t.Cleanup(func() {
		t.Log("Removing container")
		db.RunningContainer.Close()
	})
}

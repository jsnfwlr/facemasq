//go:build database

package devices

import (
	"testing"

	"facemasq/lib/db"
)

func TestGetTrendData(t *testing.T) {
	db.DBEngine = "mysql"
	container, err := db.ConnectToTest()
	if err != nil {
		t.Fatal(err)
	}
	defer container.Close()
	data, err := getTrendData()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", data)
	err = container.Close()
	if err != nil {
		t.Errorf("Error closing container: %v", err)
	}
}

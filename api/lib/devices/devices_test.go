package devices

import (
	"testing"

	"facemasq/lib/db"
)

func TestGetDeviceChildren(t *testing.T) {
	db.DBEngine = "mysql"
	container, err := db.ConnectToTest()
	if err != nil {
		t.Fatalf("Could not connect to DB: %v", err)
	}
	defer container.Close()

	children, err := GetDeviceChildren(1)
	if err != nil {
		t.Fatalf("Could not get children of device %d: %v", 1, err)
	}
	if children.Interfaces != "1" {
		t.Error("Device 1 should have Interface 1")
	}
	if children.Addresses != "1" {
		t.Error("Device 1 should have Address 1")
	}

	children, err = GetDeviceChildren(2)
	if err != nil {
		t.Fatalf("Could not get children of device %d: %v", 2, err)
	}
	if children.Interfaces != "2" {
		t.Error("Device 2 should have Interface 2")
	}
	if children.Addresses != "2" {
		t.Error("Device 2 should have Address 2")
	}

	children, err = GetDeviceChildren(3)
	if err != nil {
		t.Fatalf("Could not get children of device %d: %v", 3, err)
	}
	if children.Interfaces != "3,4,4" {
		t.Error("Device 3 should have Interfaces 3,4,4")
	}
	if children.Addresses != "3,4,5" {
		t.Error("Device 3 should have Address 3,4,5")
	}

	children, err = GetDeviceChildren(4)
	if err != nil {
		t.Fatalf("Could not get children of device %d: %v", 4, err)
	}
	if children.Interfaces != "" {
		t.Error("Device 4 should have no Interfaces")
	}
	if children.Addresses != "" {
		t.Error("Device 4 should have no Address")
	}

	err = DeleteDevice(3)
	if err != nil {
		t.Errorf("%+v", err)
	}

	children, err = GetDeviceChildren(3)
	if err != nil {
		t.Fatalf("Could not get children of device %d: %v", 3, err)
	}
	if children.Interfaces != "" {
		t.Error("Device 3 should have no Interfaces")
	}
	if children.Addresses != "" {
		t.Error("Device 3 should have no Address")
	}
	err = container.Close()
	if err != nil {
		t.Errorf("Error closing container: %v", err)
	}
}

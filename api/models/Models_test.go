package Models

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jsnfwlr/facemasq/api/lib/db"
	"github.com/jsnfwlr/facemasq/api/models/Addresses"
	"github.com/jsnfwlr/facemasq/api/models/Architectures"
	"github.com/jsnfwlr/facemasq/api/models/Categories"
	"github.com/jsnfwlr/facemasq/api/models/DeviceTypes"
	"github.com/jsnfwlr/facemasq/api/models/Devices"
	"github.com/jsnfwlr/facemasq/api/models/History"
	"github.com/jsnfwlr/facemasq/api/models/Hostnames"
	"github.com/jsnfwlr/facemasq/api/models/InterfaceTypes"
	"github.com/jsnfwlr/facemasq/api/models/Locations"
	"github.com/jsnfwlr/facemasq/api/models/Maintainers"
	"github.com/jsnfwlr/facemasq/api/models/Meta"
	"github.com/jsnfwlr/facemasq/api/models/Netfaces"
	"github.com/jsnfwlr/facemasq/api/models/OperatingSystems"
	"github.com/jsnfwlr/facemasq/api/models/Scans"
	"github.com/jsnfwlr/facemasq/api/models/Status"
	"github.com/jsnfwlr/facemasq/api/models/Users"
	"github.com/jsnfwlr/facemasq/api/models/VLANs"
	"github.com/volatiletech/null"
)

func TestDB(test *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		test.Error(err)
	}
	folder := fmt.Sprintf("%sdata%c", strings.Replace(cwd, "api/models", "", -1), os.PathSeparator)
	file := "test.sqlite"
	os.Remove(fmt.Sprintf("%s%s", folder, file))
	err = db.Connect(folder, file)
	if err != nil {
		test.Error(err)
	}
}

func TestHistory(test *testing.T) {
	_, err := History.Get()
	if err != nil {
		test.Error(err)
	}
}

func TestArchitectures(test *testing.T) {
	record := Architectures.Model{
		Label:    "Test",
		BitSpace: 64,
		Notes:    null.StringFrom("Testing"),
	}

	_, err := Architectures.Get()
	if err != nil {
		test.Error(err)
	}

	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.Notes = null.StringFrom("Will be deleted")
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}

}

func TestUsers(test *testing.T) {
	_, err := Users.Get()
	if err != nil {
		test.Error(err)
	}

	record := Users.Model{
		Label:           "Test",
		Username:        null.StringFrom("Guest"),
		NewPassword:     null.StringFrom("Password"),
		CanAuthenticate: true,
		AccessLevel:     0,
		Notes:           null.StringFrom("Testing"),
		IsLocked:        false,
	}
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.Notes = null.StringFrom("Will be deleted")
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	rows, err := Users.Get()
	if err != nil {
		test.Error(err)
	}
	for i := range rows {
		fmt.Printf("%+v\n", rows[i])
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}
}

func TestNetfaces(test *testing.T) {

	_, err := Netfaces.Get()
	if err != nil {
		test.Error(err)
	}

	lastSeen := time.Now()
	record := Netfaces.Model{
		MAC:             "00:00:00:00:00:00:00:00",
		IsPrimary:       true,
		IsVirtual:       false,
		InterfaceTypeID: 1,
		LastSeen:        lastSeen,
		VLANID:          1,
		DeviceID:        1,
		Label:           null.StringFrom("Test"),
		Notes:           null.StringFrom("Testing"),
	}

	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.Notes = null.StringFrom("Will be deleted")
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	recover, err := Netfaces.Get()
	if err != nil {
		test.Error(err)
	}

	if recover[0].LastSeen.Format("2006-01-02 15:04:05.999") != lastSeen.Format("2006-01-02 15:04:05.999") {
		test.Errorf("LastSeen mismatch %v vs %v", recover[0].LastSeen, lastSeen)
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}
}

func TestCategories(test *testing.T) {
	record := Categories.Model{
		Label: "Test",
		Icon:  "Eye",
		Notes: null.StringFrom("Testing"),
	}

	_, err := Categories.Get()
	if err != nil {
		test.Error(err)
	}

	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.Notes = null.StringFrom("Will be deleted")
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}

}

func TestScans(test *testing.T) {

	_, err := Scans.Get()
	if err != nil {
		test.Error(err)
	}
}

func TestVLANs(test *testing.T) {

	_, err := VLANs.Get()
	if err != nil {
		test.Error(err)
	}

	record := VLANs.Model{
		Label:    "",
		IPv4Mask: "129.168.1.0/24",
		IPv6Mask: "ffff::1",
		Notes:    null.StringFrom("Test"),
	}

	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.Notes = null.StringFrom("To be deleted")
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}

}

func TestDeviceTypes(test *testing.T) {

	_, err := DeviceTypes.Get()
	if err != nil {
		test.Error(err)
	}

	record := DeviceTypes.Model{
		Label: "Test",
		Icon:  "Eye",
		Notes: null.StringFrom("Testing"),
	}
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.Notes = null.StringFrom("Will be deleted")
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}
}

func TestLocations(test *testing.T) {

	_, err := Locations.Get()
	if err != nil {
		test.Error(err)
	}
	record := Locations.Model{
		Label: "Test",
		Notes: null.StringFrom("Testing"),
	}
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.Notes = null.StringFrom("Will be deleted")
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}
}

func TestHostnames(test *testing.T) {

	_, err := Hostnames.Get()
	if err != nil {
		test.Error(err)
	}
	record := Hostnames.Model{
		Hostname:  "Test",
		Notes:     null.StringFrom("Testing"),
		IsDNS:     true,
		IsSelfSet: true,
		AddressID: 1,
	}
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.Notes = null.StringFrom("Will be deleted")
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}
}

func TestOperatingSystems(test *testing.T) {

	_, err := OperatingSystems.Get()
	if err != nil {
		test.Error(err)
	}

	record := OperatingSystems.Model{
		Vendor:       "Test",
		Family:       "Test",
		Name:         "Test",
		Version:      "Test",
		IsOpenSource: false,
		IsServer:     true,
		Notes:        null.StringFrom("Testing"),
	}
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.Notes = null.StringFrom("Will be deleted")
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}
}

func TestInterfaceTypes(test *testing.T) {

	_, err := InterfaceTypes.Get()
	if err != nil {
		test.Error(err)
	}
	record := InterfaceTypes.Model{
		Label: "Test",
		Icon:  "Eye",
		Notes: null.StringFrom("Testing"),
	}
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.Notes = null.StringFrom("Will be deleted")
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}
}

func TestStatus(test *testing.T) {

	_, err := Status.Get()
	if err != nil {
		test.Error(err)
	}

	record := Status.Model{
		Label: "Test",
		Icon:  "Eye",
		Notes: null.StringFrom("Testing"),
	}
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.Notes = null.StringFrom("Will be deleted")
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}
}

func TestAddresses(test *testing.T) {
	_, err := Addresses.Get()
	if err != nil {
		test.Error(err)
	}

	record := Addresses.Model{
		IPv4:        null.StringFrom("192.168.5.1"),
		IPv6:        null.NewString("", false),
		IsPrimary:   true,
		IsVirtual:   false,
		IsReserved:  false,
		LastSeen:    time.Now(),
		Label:       null.StringFrom("Test"),
		Notes:       null.StringFrom("This should be deleted"),
		InterfaceID: 1,
	}

	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.IPv4 = null.StringFrom("192.168.5.2")
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}

}

func TestMaintainers(test *testing.T) {

	_, err := Maintainers.Get()
	if err != nil {
		test.Error(err)
	}

	record := Maintainers.Model{
		Label:      "Test",
		IsInternal: true,
		Notes:      null.StringFrom("Will Be Deleted"),
	}
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.IsInternal = false
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}
}

func TestDevices(test *testing.T) {

	_, err := Devices.Get()
	if err != nil {
		test.Error(err)
	}

	record := Devices.Model{
		Label:     null.StringFrom("Testing"),
		Notes:     null.StringFrom("Testing"),
		FirstSeen: null.StringFrom(time.Now().Format("2006-01-02 15:04:05")),
	}
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.Notes = null.StringFrom("Will be deleted")
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}
}

func TestMeta(test *testing.T) {
	var records Meta.Models
	sql := `SELECT Name, Value FROM Meta WHERE UserID IS NULL`
	err := db.Conn.Select(&records, sql)
	if err != nil {
		test.Error(err)
	}
	record := Meta.Model{
		Name:  "Test",
		Value: "Testing",
	}
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	record.Value = "Will be deleted"
	err = record.Save()
	if err != nil {
		test.Error(err)
	}

	err = record.Delete()
	if err != nil {
		test.Error(err)
	}
}

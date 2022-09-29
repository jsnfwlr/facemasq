package devices

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"facemasq/lib/db"
	helper "facemasq/lib/devices"
	"facemasq/lib/formats"
	"facemasq/models"
)

var DefaultConnTime = time.Duration(-30) * time.Minute

type Investigation struct {
	AddressID    int64
	Connectivity []models.Connections
}

func GetActive(out http.ResponseWriter, in *http.Request) {
	queries := helper.DeviceQueries{
		Devices:   `SELECT * FROM devices;`,
		Netfaces:  `SELECT * FROM interfaces ORDER BY is_primary DESC, is_virtual ASC;`,
		Addresses: `SELECT * FROM addresses WHERE last_seen = (SELECT time FROM scans ORDER BY time DESC LIMIT 1) ORDER BY is_primary DESC, is_virtual ASC;`,
		Hostnames: `SELECT * FROM hostnames;`,
	}

	activeDevices, err := helper.GetDevices(queries, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"), true)
	if err != nil {
		log.Panicf("Error: %v", err)
		http.Error(out, fmt.Sprintf("Unable to retrieve data: %v", err), http.StatusInternalServerError)
		return
	}

	formats.WriteJSONResponse(activeDevices, out, in)
}

func GetAll(out http.ResponseWriter, in *http.Request) {
	queries := helper.DeviceQueries{
		Devices:   `SELECT * FROM devices;`,
		Netfaces:  `SELECT * FROM interfaces ORDER BY is_primary DESC, is_virtual ASC;`,
		Addresses: `SELECT * FROM addresses ORDER BY interface_id ASC, is_primary DESC, last_seen DESC, is_virtual ASC;`,
		Hostnames: `SELECT * FROM hostnames;`,
	}
	allDevices, err := helper.GetDevices(queries, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"), true)
	if err != nil {
		log.Panicf("Error: %v", err)
		http.Error(out, fmt.Sprintf("Unable to retrieve data: %v", err), http.StatusInternalServerError)
		return
	}

	formats.WriteJSONResponse(allDevices, out, in)
}

func GetUnknown(out http.ResponseWriter, in *http.Request) {
	queries := helper.DeviceQueries{
		Devices:   `SELECT * FROM devices WHERE is_tracked = 0;`,
		Netfaces:  `SELECT * FROM interfaces ORDER BY is_primary DESC, is_virtual ASC;`,
		Addresses: `SELECT * FROM addresses ORDER BY is_primary DESC, is_virtual ASC;`,
		Hostnames: `SELECT * FROM hostnames;`,
	}

	unknownDevices, err := helper.GetDevices(queries, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"), true)
	if err != nil {
		log.Panicf("Error: %v", err)
		http.Error(out, fmt.Sprintf("Unable to retrieve data: %v", err), http.StatusInternalServerError)
		return
	}

	formats.WriteJSONResponse(unknownDevices, out, in)
}

func SaveDevice(out http.ResponseWriter, in *http.Request) {
	var input models.Device
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		log.Printf("Unable to parse Device: %v", err)
		http.Error(out, "Unable to parse Device", http.StatusInternalServerError)
		return
	}
	if input.FirstSeen.Format("2006-01-02") == "0001-01-01" {
		input.FirstSeen = time.Now()
	}
	log.Printf("!! %v", input)
	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where("id = ?", input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("Unable to save Device: %v", err)
		http.Error(out, "Unable to save Device", http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func SaveInterface(out http.ResponseWriter, in *http.Request) {
	var input models.Interface
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		log.Printf("Unable to parse Inteface: %v", err)
		http.Error(out, "Unable to parse Inteface", http.StatusInternalServerError)
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where("id = ?", input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("Unable to save Inteface: %v", err)
		http.Error(out, "Unable to save Inteface", http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func SaveAddress(out http.ResponseWriter, in *http.Request) {
	var input models.Address
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		log.Printf("Unable to parse Address: %v", err)
		http.Error(out, "Unable to parse Address", http.StatusInternalServerError)
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("Unable to save Address: %v", err)
		http.Error(out, "Unable to save Address", http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func SaveHostname(out http.ResponseWriter, in *http.Request) {
	var input models.Hostname
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		log.Printf("Unable to parse Hostname: %v", err)
		http.Error(out, "Unable to parse Hostname", http.StatusInternalServerError)
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("Unable to save Hostname: %v", err)
		http.Error(out, "Unable to save Hostname", http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func InvestigateAddresses(out http.ResponseWriter, in *http.Request) {
	var input []int64
	var investigations []Investigation
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		log.Printf("Unable to parse Address IDs: %v\n", err)
		http.Error(out, "Unable to parse Address IDs", http.StatusInternalServerError)
		return
	}
	for _, address := range input {
		investigation := Investigation{
			AddressID: address,
		}
		investigation.Connectivity, err = helper.GetSpecificAddressConnectivityData(address, time.Now().Add(time.Duration(7*24*-1)*time.Hour).Format("2006-01-02 15:04"))
		if err != nil {
			log.Printf("Unable to get connection data: %v\n", err)
			http.Error(out, "Unable to get connection data", http.StatusInternalServerError)
			return
		}
		investigations = append(investigations, investigation)
	}
	formats.WriteJSONResponse(investigations, out, in)
}

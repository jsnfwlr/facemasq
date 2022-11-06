package devices

import (
	"fmt"
	"net/http"
	"time"

	helper "facemasq/lib/devices"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/models"

	"github.com/gorilla/websocket"
)

var DefaultConnTime = time.Duration(-60) * time.Minute

type Investigation struct {
	AddressID    int64
	Connectivity []models.Connection
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func InvestigateAddresses(out http.ResponseWriter, in *http.Request) {
	var input []int64
	var investigations []Investigation
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		logging.Error("Unable to parse Address IDs: %v\n", err)
		http.Error(out, "Unable to parse Address IDs", http.StatusInternalServerError)
		return
	}
	for _, address := range input {
		investigation := Investigation{
			AddressID: address,
		}
		investigation.Connectivity, err = helper.GetSpecificAddressConnectivityData(address, time.Now().Add(time.Duration(7*24*-1)*time.Hour).Format("2006-01-02 15:04"))
		if err != nil {
			logging.Error("Unable to get connection data: %v\n", err)
			http.Error(out, "Unable to get connection data", http.StatusInternalServerError)
			return
		}
		investigations = append(investigations, investigation)
	}
	formats.WriteJSONResponse(investigations, out, in)
}

func GetActive(out http.ResponseWriter, in *http.Request) {
	queries := helper.DeviceQueries{
		Devices:   `SELECT * FROM devices;`,
		Netfaces:  `SELECT * FROM interfaces ORDER BY is_primary DESC, is_virtual ASC;`,
		Addresses: `SELECT * FROM addresses WHERE last_seen = (SELECT time FROM scans ORDER BY time DESC LIMIT 1) ORDER BY is_primary DESC, is_virtual ASC;`,
		Hostnames: `SELECT * FROM hostnames;`,
		// Devices:   `SELECT * FROM devices;`,
		// Netfaces:  `SELECT * FROM interfaces ORDER BY is_primary DESC, is_virtual ASC;`,
		// Addresses: `SELECT * FROM addresses WHERE last_seen = (SELECT time FROM scans ORDER BY time DESC LIMIT 1) ORDER BY is_primary DESC, is_virtual ASC;`,
		// Hostnames: `SELECT * FROM hostnames;`,
	}

	activeDevices, err := helper.GetDevices(queries, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"), true)
	if err != nil {
		logging.Panic("Error: %v", err)
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
		logging.Panic("Error: %v", err)
		http.Error(out, fmt.Sprintf("Unable to retrieve data: %v", err), http.StatusInternalServerError)
		return
	}

	formats.WriteJSONResponse(allDevices, out, in)
}

func GetUnknown(out http.ResponseWriter, in *http.Request) {
	queries := helper.DeviceQueries{
		Devices:   `SELECT * FROM devices WHERE status_id = 1;`,
		Netfaces:  `SELECT * FROM interfaces ORDER BY is_primary DESC, is_virtual ASC;`,
		Addresses: `SELECT * FROM addresses ORDER BY is_primary DESC, is_virtual ASC;`,
		Hostnames: `SELECT * FROM hostnames;`,
	}

	unknownDevices, err := helper.GetDevices(queries, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"), true)
	if err != nil {
		logging.Panic("Error: %v", err)
		http.Error(out, fmt.Sprintf("Unable to retrieve data: %v", err), http.StatusInternalServerError)
		return
	}

	formats.WriteJSONResponse(unknownDevices, out, in)
}

func GetRecentChanges(out http.ResponseWriter, in *http.Request) {
	socket, err := upgrader.Upgrade(out, in, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			logging.Error(err)
		}
		return
	}

	lastMod, _ := time.Parse("2006-01-02 15:04:05", in.FormValue("lastMod"))

	go helper.Writer(socket, lastMod)
	helper.Reader(socket)
}

/*
(function() {
		var data = document.querySelector(".flex.items-center.py-0.px-3.bg-gray-100");
		var conn = new WebSocket("ws://192.168.0.41:6135/ws/records/changed?lastMod=2022-09-25 11:40:00");
		conn.onclose = function(evt) {
				data.textContent = 'Connection closed';
		}
		conn.onmessage = function(evt) {
				console.log('file updated');
				data.textContent = evt.data;
		}
})();
*/

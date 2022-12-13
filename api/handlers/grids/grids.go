package grids

import (
	"net/http"
	"time"

	helper "facemasq/lib/devices"
	"facemasq/lib/formats"
	"facemasq/models"

	"github.com/gorilla/websocket"
	"github.com/uptrace/bunrouter"
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

func GetDeviceDosier(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var input []int64
	var investigations []Investigation
	err = formats.ReadJSONBody(in, &input)
	if err != nil {
		return
	}
	for _, address := range input {
		investigation := Investigation{
			AddressID: address,
		}
		investigation.Connectivity, err = helper.GetSpecificAddressConnectivityData(address, time.Now().Add(time.Duration(7*24*-1)*time.Hour).Format("2006-01-02 15:04"))
		if err != nil {
			return
		}
		investigations = append(investigations, investigation)
	}
	formats.WriteJSONResponse(investigations, out, in)

	return
}

func GetActiveDevices(out http.ResponseWriter, in bunrouter.Request) (err error) {
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
		return
	}

	formats.WriteJSONResponse(activeDevices, out, in)
	return
}

func GetAllDevices(out http.ResponseWriter, in bunrouter.Request) (err error) {
	queries := helper.DeviceQueries{
		Devices:   `SELECT * FROM devices;`,
		Netfaces:  `SELECT * FROM interfaces ORDER BY is_primary DESC, is_virtual ASC;`,
		Addresses: `SELECT * FROM addresses ORDER BY interface_id ASC, is_primary DESC, last_seen DESC, is_virtual ASC;`,
		Hostnames: `SELECT * FROM hostnames;`,
	}
	allDevices, err := helper.GetDevices(queries, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"), true)
	if err != nil {
		return
	}

	formats.WriteJSONResponse(allDevices, out, in)
	return
}

func GetUnknownDevices(out http.ResponseWriter, in bunrouter.Request) (err error) {
	queries := helper.DeviceQueries{
		Devices:   `SELECT * FROM devices WHERE status_id = 1;`,
		Netfaces:  `SELECT * FROM interfaces ORDER BY is_primary DESC, is_virtual ASC;`,
		Addresses: `SELECT * FROM addresses ORDER BY is_primary DESC, is_virtual ASC;`,
		Hostnames: `SELECT * FROM hostnames;`,
	}

	unknownDevices, err := helper.GetDevices(queries, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"), true)
	if err != nil {
		return
	}

	formats.WriteJSONResponse(unknownDevices, out, in)
	return
}

// func GetRecentlyChangedDevices(out http.ResponseWriter, in bunrouter.Request) (err error) {
// 	socket, err := upgrader.Upgrade(out, in, nil)
// 	if err != nil {
// 		if _, ok := err.(websocket.HandshakeError); !ok {
// 			logging.Error(err)
// 		}
// 		return
// 	}

// 	lastMod, _ := time.Parse("2006-01-02 15:04:05", in.FormValue("lastMod"))

// 	go helper.Writer(socket, lastMod)
// 	helper.Reader(socket)
// }

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

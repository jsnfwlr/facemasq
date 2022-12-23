package grids

import (
	"encoding/json"
	"net/http"
	"time"

	helper "facemasq/lib/devices"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/models"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/uptrace/bunrouter"
)

var DefaultConnTime = time.Duration(-60) * time.Minute

type Investigation struct {
	AddressID    int64
	Connectivity []models.Connection
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
	activeDevices, err := helper.GetDevices(false, true, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"), true)
	if err != nil {
		return
	}

	formats.WriteJSONResponse(activeDevices, out, in)
	return
}

func GetAllDevices(out http.ResponseWriter, in bunrouter.Request) (err error) {
	allDevices, err := helper.GetDevices(false, false, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"), true)
	if err != nil {
		return
	}

	formats.WriteJSONResponse(allDevices, out, in)
	return
}

func GetUnknownDevices(out http.ResponseWriter, in bunrouter.Request) (err error) {
	unknownDevices, err := helper.GetDevices(true, false, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"), true)
	if err != nil {
		return
	}

	formats.WriteJSONResponse(unknownDevices, out, in)
	return
}

func GetRecentlyChangedDevices(out http.ResponseWriter, in bunrouter.Request) (err error) {
	socket, _, _, err := ws.UpgradeHTTP(in.Request, out)
	if err != nil {
		logging.Error(err)
		return
	}
	// @TODO: move to channels so events can send data
	go func() {
		defer socket.Close()
		for {
			msg, op, err := wsutil.ReadClientData(socket)
			if err != nil {
				logging.Error(err)
				break
			}
			lastMod, _ := time.Parse("2006-01-02 15:04:05", string(msg))
			logging.Debug(lastMod)
			changes, lastSeen, err := helper.GetChangesSince(lastMod, true)
			if err != nil {
				logging.Error(err)
			}
			reply := make(map[string]interface{})
			reply["lastMod"] = lastSeen
			reply["devices"] = changes
			msg, err = json.Marshal(reply)
			if err != nil {
				logging.Error(err)
			}
			err = wsutil.WriteServerMessage(socket, op, msg)
			if err != nil {
				logging.Error(err)
				break
			}
		}
	}()

	// lastMod, _ := time.Parse("2006-01-02 15:04:05", in.FormValue("lastMod"))

	// done := false
	// for {
	// 	switch {
	// 	case <-appDone:
	// 		done = true
	// 	case <-appUpdate:
	// 		go sendMessage(socket, helper.GetChangesSince(&lastMod, true))
	// 	}
	// 	if done {
	// 		break
	// 	}
	// }
	return
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

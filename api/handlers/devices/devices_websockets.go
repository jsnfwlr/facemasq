package devices

import (
	"net/http"
	"time"

	"facemasq/lib/devices"
	"facemasq/lib/logging"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func GetRecentChanges(out http.ResponseWriter, in *http.Request) {
	socket, err := upgrader.Upgrade(out, in, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			logging.Errorln(err)
		}
		return
	}

	lastMod, _ := time.Parse("2006-01-02 15:04:05", in.FormValue("lastMod"))

	go devices.Writer(socket, lastMod)
	devices.Reader(socket)
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

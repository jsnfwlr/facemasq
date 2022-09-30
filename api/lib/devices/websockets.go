package devices

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	WriteWait  = 10 * time.Second
	PongWait   = 60 * time.Second
	PingPeriod = (PongWait * 9) / 10
)

var ScanPeriod = 10 * time.Second

func Reader(socket *websocket.Conn) {
	defer socket.Close()
	socket.SetReadLimit(512)
	socket.SetReadDeadline(time.Now().Add(PongWait))
	socket.SetPongHandler(func(string) error { socket.SetReadDeadline(time.Now().Add(PongWait)); return nil })
	for {
		_, _, err := socket.ReadMessage()
		if err != nil {
			break
		}
	}
}

func Writer(socket *websocket.Conn, lastMod time.Time) {
	// lastError := ""
	pingTicker := time.NewTicker(PingPeriod)
	fileTicker := time.NewTicker(ScanPeriod)

	defer func() {
		pingTicker.Stop()
		fileTicker.Stop()
		socket.Close()
	}()
	for {
		select {
		case <-fileTicker.C:
			var output []byte
			// var allDevices []models.Device
			// var err error

			/*
				allDevices, lastMod, err = GetChangesSince(lastMod, true)
				if err != nil {
					if s := err.Error(); s != lastError {
						lastError = s
						output = []byte(lastError)
					}
				} else {
					lastError = ""
				}
				if len(allDevices) > 0 {
					output, err = json.Marshal(allDevices)
					if err != nil {
						if s := err.Error(); s != lastError {
							lastError = s
							output = []byte(lastError)
						}
					} else {
						lastError = ""
					}
				}
			*/
			output = []byte(lastMod.Format("15:04:05 2006-01-02"))

			if output != nil {
				socket.SetWriteDeadline(time.Now().Add(WriteWait))
				if err := socket.WriteMessage(websocket.TextMessage, output); err != nil {
					return
				}
			}
		case <-pingTicker.C:
			socket.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := socket.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

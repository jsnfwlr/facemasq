package main

import (
	"bytes"
	"encoding/json"
	"facemasq/lib/apiclient"
	"facemasq/lib/events"
	"facemasq/lib/extensions"
	"facemasq/lib/logging"
	"os"
)

var GotifyURL, GotifyKey string
var EnableGotify bool

func init() {
	EnableGotify = true
	GotifyURL = os.Getenv("GOTIFY_URL")

	GotifyKey = os.Getenv("GOTIFY_KEY")
	if GotifyURL == "" || GotifyKey == "" {
		EnableGotify = false
	}

}

type MsgBody struct {
	Message  string
	Priority int64
	Title    string
}

func LoadExtension(manager *extensions.Manager) (err error) {

	manager.RegisterListeners([]extensions.Listener{
		{Kind: `notification:send:.*`, Listener: SendMessage},
	})
	return
}

func SendMessage(event events.Event) {

	payload, err := json.Marshal(event.Payload)
	if err != nil {
		logging.Error("Error sending notification via gotify: %v", err)
		return
	}

	var payLoad interface{}

	err = json.Unmarshal(payload, &payLoad)
	if err != nil {
		logging.Error("Error sending notification via gotify: %v", err)
		return
	}
	title := payLoad.(map[string]interface{})["title"].(string)
	message := payLoad.(map[string]interface{})["message"].(string)
	priority := int64(payLoad.(map[string]interface{})["priority"].(float64))
	// var response *http.Response
	if EnableGotify {
		var bodyJSON []byte
		body := MsgBody{
			Message:  message,
			Priority: priority,
			Title:    title,
		}

		headers := make(apiclient.Headers)

		headers["X-Gotify-Key"] = GotifyKey
		headers["Content-Type"] = "application/json"
		headers["accept"] = "application/json"

		bodyJSON, err = json.Marshal(body)
		if err != nil {
			logging.Error("Error sending notification via gotify: %v", err)
			return
		}

		client := apiclient.Prepare()

		request := apiclient.Request{
			URL:      GotifyURL,
			Method:   "POST",
			Headers:  headers,
			BodyJSON: bytes.NewReader(bodyJSON),
		}
		_, err = client.Do(&request)
		if err != nil {
			logging.Error("Error sending notification via gotify: %v", err)
			return
		}
	}
}

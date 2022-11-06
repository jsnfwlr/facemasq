package main

import (
	"bytes"
	"encoding/json"
	"facemasq/lib/apiclient"
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

func SendMessage(title, message string, priority int64) (err error) {
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
			return
		}
	}
	return
}

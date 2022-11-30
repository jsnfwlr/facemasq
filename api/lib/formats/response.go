package formats

import (
	"bytes"
	"io"
	"net/http"

	"facemasq/lib/logging"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func WriteTextResponse(payload string, response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/plain")
	response.Write([]byte(payload))
}

func WriteJSONResponse(payload interface{}, response http.ResponseWriter, request *http.Request) {
	json, _ := json.Marshal(payload)
	response.Header().Set("Content-Type", "application/json")
	response.Write(json)
}

func ReadJSONBody(request *http.Request, target interface{}) (err error) {
	var body []byte
	body, err = io.ReadAll(request.Body)
	if err != nil {
		return
	}
	request.Body = io.NopCloser(bytes.NewBuffer(body))
	logging.Debug("Request Body Contents: %v", string(body))
	err = json.Unmarshal(body, target)
	return
}

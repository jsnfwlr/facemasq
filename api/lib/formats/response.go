package formats

import (
	"bytes"
	"io"
	"net/http"

	"facemasq/lib/logging"

	jsoniter "github.com/json-iterator/go"
	"github.com/uptrace/bunrouter"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func WriteTextResponse(payload string, out http.ResponseWriter, request bunrouter.Request) {
	out.Header().Set("Content-Type", "text/plain")
	out.Write([]byte(payload))
}

func WriteJSONResponse(payload interface{}, out http.ResponseWriter, in bunrouter.Request) {
	bunrouter.JSON(out, payload)
	// json, _ := json.Marshal(payload)

	// response.Header().Set("Content-Type", "application/json")
	// response.Write(json)
}

func ReadJSONBody(in bunrouter.Request, target interface{}) (err error) {
	var body []byte
	body, err = io.ReadAll(in.Body)
	if err != nil {
		return
	}
	in.Body = io.NopCloser(bytes.NewBuffer(body))
	logging.Debug("Request Body Contents: %v", string(body))
	err = json.Unmarshal(body, target)
	return
}

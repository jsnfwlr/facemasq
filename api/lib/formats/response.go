package formats

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func PublishText(payload string, response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/plain")
	response.Write([]byte(payload))
}

func PublishJSON(payload interface{}, response http.ResponseWriter, request *http.Request) {
	json, _ := json.Marshal(payload)
	response.Header().Set("Content-Type", "application/json")
	response.Write(json)
}

func ReadJSON(request *http.Request, target interface{}) (err error) {
	var body []byte
	body, err = ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	log.Printf("%v\n", string(body))
	err = json.Unmarshal(body, target)
	return
}

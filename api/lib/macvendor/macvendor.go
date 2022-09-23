package macvendor

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const uri = "https://api.macvendors.com/%s"

func Lookup(mac string) (vendor string, err error) {
	var response *http.Response
	response, err = http.Get(fmt.Sprintf(uri, mac))
	if err != nil {
		return
	}
	buf := new(strings.Builder)
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		return
	}
	vendor = buf.String()
	return
}

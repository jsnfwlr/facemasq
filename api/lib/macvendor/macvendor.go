package macvendor

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const uri = "https://api.macvendors.com/%s"

var TooManyRequests = false

func Lookup(mac string) (vendor string, err error) {
	if !TooManyRequests {
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
		if strings.Contains(vendor, "errors") {
			if strings.Contains(vendor, "Too Many Requests") {
				TooManyRequests = true
			}
			err = fmt.Errorf("could not determine vendor of `%s`: %v", mac, vendor)
			vendor = ""
		}
		return
	}
	err = fmt.Errorf("too many vendor requests")
	return

}

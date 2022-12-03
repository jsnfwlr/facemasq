package macvendor

import (
	"facemasq/lib/apiclient"
	"facemasq/lib/logging"
	"fmt"
	"time"
)

const uri = "https://api.macvendors.com/%s"

var TooManyRequests = time.Now()
var LockTime = 2.5

func Lookup(mac string) (vendor string, err error) {
	var response *apiclient.Response
	if TooManyRequests.After(time.Now()) {
		err = fmt.Errorf("could not determine vendor of `%s`: Too Many Requests - locked for %f seconds", mac, LockTime)
		return
	}
	client := apiclient.Prepare()

	request := apiclient.Request{
		URL:    fmt.Sprintf(uri, mac),
		Method: "GET",
	}
	response, err = client.Do(&request)
	if err != nil {
		if err.Error() == "Too Many Requests" {
			TooManyRequests = time.Now().Add(time.Duration(LockTime*1000) * time.Millisecond)
			logging.Info("Temporarily disabling MAC vendor queries for %f seconds due to rate limits", LockTime)
			err = fmt.Errorf("could not determine vendor of `%s`: Too Many Requests - locked for %f seconds", mac, LockTime)
			return
		}
		err = fmt.Errorf("could not determine vendor of `%s`: %v", mac, err)
		return
	}
	vendor = response.Body
	return
}

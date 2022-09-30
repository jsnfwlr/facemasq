package network

import (
	"os"

	"facemasq/lib/logging"
)

var Target string

func init() {
	Target = os.Getenv("NETMASK")
}

func ShowNetworkSummary() (err error) {
	var (
		addresses                []string
		network, broadcast, mask string
	)
	addresses, network, broadcast, mask, err = getIPRange(Target)
	if err != nil {
		return
	}
	logging.Processf("Network: %s, Broadcast: %s, Mask: %s, Addresses: %d\n", network, broadcast, mask, len(addresses))
	return
}

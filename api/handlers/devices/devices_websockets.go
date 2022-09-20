package devices

import (
	"net/http"
	"time"

	helper "facemasq/lib/devices"
	"facemasq/lib/formats"
)

func GetActiveWS(out http.ResponseWriter, in *http.Request) {
	queries := helper.DeviceQueries{
		Devices:   `SELECT * FROM Devices;`,
		Netfaces:  `SELECT * FROM Interfaces ORDER BY IsPrimary DESC, IsVirtual ASC;`,
		Addresses: `SELECT * FROM Addresses WHERE LastSeen = (SELECT Time FROM Scans ORDER BY Time DESC LIMIT 1) ORDER BY IsPrimary DESC, IsVirtual ASC;`,
		Hostnames: `SELECT * FROM Hostnames;`,
	}

	activeDevices, err := helper.GetDevices(queries, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"), true)
	if err != nil {
		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
	}

	formats.PublishJSON(activeDevices, out, in)
}

func GetAllWS(out http.ResponseWriter, in *http.Request) {
	queries := helper.DeviceQueries{
		Devices:   `SELECT * FROM Devices;`,
		Netfaces:  `SELECT * FROM Interfaces ORDER BY IsPrimary DESC, IsVirtual ASC;`,
		Addresses: `SELECT * FROM Addresses ORDER BY InterfaceID ASC, IsPrimary DESC, LastSeen DESC, IsVirtual ASC;`,
		Hostnames: `SELECT * FROM Hostnames;`,
	}
	allDevices, err := helper.GetDevices(queries, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"), true)
	if err != nil {
		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
	}

	formats.PublishJSON(allDevices, out, in)
}

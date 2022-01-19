package devices

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jsnfwlr/facemasq/api/lib/db"
	"github.com/jsnfwlr/facemasq/api/lib/formats"
	"github.com/jsnfwlr/facemasq/api/models/Addresses"
	"github.com/jsnfwlr/facemasq/api/models/Connections"
	"github.com/jsnfwlr/facemasq/api/models/Devices"
	"github.com/jsnfwlr/facemasq/api/models/Hostnames"
	"github.com/jsnfwlr/facemasq/api/models/Netfaces"
	"github.com/volatiletech/null"
)

var DefaultConnTime = time.Duration(-30) * time.Minute

type DeviceQueries struct {
	Devices   string
	Netfaces  string
	Addresses string
	Hostnames string
}
type Connectivity []Connections.Model

func (a Connectivity) Less(i, j int) bool { return a[i].Time.Before(a[j].Time) }
func (a Connectivity) Len() int           { return len(a) }
func (a Connectivity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func getConnectivityData(addressID int64, connTime string) (connections []Connections.Model, err error) {
	sql := `SELECT Scans.Time, CASE WHEN AddressID IS NULL THEN false ELSE true END as State FROM Scans LEFT JOIN History ON History.ScanID = Scans.ID AND History.AddressID = ? WHERE Scans.Time > ? ORDER BY Scans.ID DESC;`
	err = db.Conn.Select(&connections, sql, addressID, connTime)
	sort.Sort(Connectivity(connections))
	return
}

func getDevices(queries DeviceQueries, connTime string) (matchedDevices []Devices.Model, err error) {
	var devices []Devices.Model
	var netfaces []Netfaces.Model
	var addresses []Addresses.Model
	var hostnames []Hostnames.Model

	err = db.Conn.Select(&devices, queries.Devices)
	if err != nil {
		return
	}

	err = db.Conn.Select(&netfaces, queries.Netfaces)
	if err != nil {
		return
	}

	err = db.Conn.Select(&addresses, queries.Addresses)
	if err != nil {
		return
	}

	err = db.Conn.Select(&hostnames, queries.Hostnames)
	if err != nil {
		return
	}

	for a := range addresses {
		if addresses[a].IPv4.Valid {
			sections := strings.Split(addresses[a].IPv4.String, ".")

			for s := range sections {
				var ipnum int
				ipnum, _ = strconv.Atoi(sections[s])
				addresses[a].SortOrder += fmt.Sprintf("%03d", ipnum)
			}
		}
		addresses[a].Connectivity, err = getConnectivityData(addresses[a].ID, connTime)
		if err != nil {
			return
		}
	}

	for d := range devices {

		for n := range netfaces {
			if devices[d].ID == netfaces[n].DeviceID {
				for a := range addresses {
					if netfaces[n].ID == addresses[a].InterfaceID {
						for h := range hostnames {
							if addresses[a].ID == hostnames[h].AddressID {
								addresses[a].Hostnames = append(addresses[a].Hostnames, hostnames[h])
							}
						}
						netfaces[n].Addresses = append(netfaces[n].Addresses, addresses[a])
						if netfaces[n].SortOrder == "" {
							netfaces[n].SortOrder = addresses[a].SortOrder
							netfaces[n].Primary.IPv4 = addresses[a].IPv4.String
							netfaces[n].Primary.IPv6 = addresses[a].IPv6.String
							netfaces[n].Primary.IsVirtualIP = addresses[a].IsVirtual
							netfaces[n].Primary.IsReservedIP = addresses[a].IsReserved
						}
					}
				}
				if len(netfaces[n].Addresses) > 0 {
					devices[d].Interfaces = append(devices[d].Interfaces, netfaces[n])
					if devices[d].SortOrder == "" {
						devices[d].SortOrder = netfaces[n].SortOrder
						devices[d].Primary = netfaces[n].Primary
						devices[d].Primary.MAC = netfaces[n].MAC
						devices[d].Primary.InterfaceTypeID = netfaces[n].InterfaceTypeID
						devices[d].Primary.VLANID = netfaces[n].VLANID
						devices[d].Primary.IsVirtualIFace = netfaces[n].IsVirtual
					}
				}
			}
		}
	}

	for d := range devices {
		if len(devices[d].Interfaces) > 0 {
			matchedDevices = append(matchedDevices, devices[d])
		}
	}
	sort.SliceStable(matchedDevices, func(i, j int) bool {
		return matchedDevices[i].SortOrder < matchedDevices[j].SortOrder
	})
	return
}

func GetActive(out http.ResponseWriter, in *http.Request) {
	queries := DeviceQueries{
		Devices:   `SELECT * FROM Devices;`,
		Netfaces:  `SELECT * FROM Interfaces ORDER BY IsPrimary DESC, IsVirtual ASC;`,
		Addresses: `SELECT * FROM Addresses WHERE LastSeen = (SELECT Time FROM Scans ORDER BY Time DESC LIMIT 1) ORDER BY IsPrimary DESC, IsVirtual ASC;`,
		Hostnames: `SELECT * FROM Hostnames;`,
	}

	activeDevices, err := getDevices(queries, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"))
	if err != nil {
		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
	}

	formats.PublishJSON(activeDevices, out, in)
}

func GetAll(out http.ResponseWriter, in *http.Request) {
	queries := DeviceQueries{
		Devices:   `SELECT * FROM Devices;`,
		Netfaces:  `SELECT * FROM Interfaces ORDER BY IsPrimary DESC, IsVirtual ASC;`,
		Addresses: `SELECT * FROM Addresses ORDER BY InterfaceID ASC, IsPrimary DESC, LastSeen DESC, IsVirtual ASC;`,
		Hostnames: `SELECT * FROM Hostnames;`,
	}
	allDevices, err := getDevices(queries, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"))
	if err != nil {
		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
	}

	formats.PublishJSON(allDevices, out, in)
}

func GetUnknown(out http.ResponseWriter, in *http.Request) {
	queries := DeviceQueries{
		Devices:   `SELECT * FROM Devices WHERE IsTracked = 0;`,
		Netfaces:  `SELECT * FROM Interfaces ORDER BY IsPrimary DESC, IsVirtual ASC;`,
		Addresses: `SELECT * FROM Addresses ORDER BY IsPrimary DESC, IsVirtual ASC;`,
		Hostnames: `SELECT * FROM Hostnames;`,
	}

	unknownDevices, err := getDevices(queries, time.Now().Add(DefaultConnTime).Format("2006-01-02 15:04"))
	if err != nil {
		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
	}

	formats.PublishJSON(unknownDevices, out, in)
}

func SaveDevice(out http.ResponseWriter, in *http.Request) {
	var input Devices.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to parse Device: %v", err)
		http.Error(out, "Unable to parse Device", http.StatusInternalServerError)
		return
	}
	if !input.FirstSeen.Valid {
		input.FirstSeen = null.NewString(time.Now().Format("2006-01-02 15:04:05"), true)
	}
	err = input.Save()
	if err != nil {
		log.Printf("Unable to save Device: %v", err)
		http.Error(out, "Unable to save Device", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}

func SaveInterface(out http.ResponseWriter, in *http.Request) {
	var input Netfaces.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to parse Inteface: %v", err)
		http.Error(out, "Unable to parse Inteface", http.StatusInternalServerError)
		return
	}
	err = input.Save()
	if err != nil {
		log.Printf("Unable to save Inteface: %v", err)
		http.Error(out, "Unable to save Inteface", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}

func SaveAddress(out http.ResponseWriter, in *http.Request) {
	var input Addresses.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to parse Address: %v", err)
		http.Error(out, "Unable to parse Address", http.StatusInternalServerError)
		return
	}
	err = input.Save()
	if err != nil {
		log.Printf("Unable to save Address: %v", err)
		http.Error(out, "Unable to save Address", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}

func SaveHostname(out http.ResponseWriter, in *http.Request) {
	var input Hostnames.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to parse Hostname: %v", err)
		http.Error(out, "Unable to parse Hostname", http.StatusInternalServerError)
		return
	}
	err = input.Save()
	if err != nil {
		log.Printf("Unable to save Hostname: %v", err)
		http.Error(out, "Unable to save Hostname", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}

type Investigation struct {
	AddressID    int64
	Connectivity []Connections.Model
}

func InvestigateAddresses(out http.ResponseWriter, in *http.Request) {
	var input []int64
	var investigations []Investigation
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to parse Address IDs: %v\n", err)
		http.Error(out, "Unable to parse Address IDs", http.StatusInternalServerError)
		return
	}
	for _, address := range input {
		investigation := Investigation{
			AddressID: address,
		}
		investigation.Connectivity, err = getConnectivityData(address, time.Now().Add(time.Duration(7*24*-1)*time.Hour).Format("2006-01-02 15:04"))
		if err != nil {
			log.Printf("Unable to get connection data: %v\n", err)
			http.Error(out, "Unable to get connection data", http.StatusInternalServerError)
			return
		}
		investigations = append(investigations, investigation)
	}
	formats.PublishJSON(investigations, out, in)
}

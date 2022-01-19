package dnsmasq

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/jsnfwlr/facemasq/api/lib/db"
	"github.com/jsnfwlr/facemasq/api/lib/files"
)

type DNS struct {
	Hostname  string `db:"Hostname"`
	IPv4      string `db:"IPv4"`
	Label     string `db:"Label"`
	SortOrder string
}

func WriteDNSConfig(out http.ResponseWriter, in *http.Request) {
	var records []DNS
	sql := `SELECT Hostname, IPv4, Devices.Label FROM Hostnames JOIN Addresses ON Hostnames.AddressID = Addresses.ID JOIN Interfaces ON Addresses.InterfaceID = Interfaces.ID JOIN Devices ON Devices.ID = Interfaces.DeviceID WHERE IsDNS = 1;`
	err := db.Conn.Select(&records, sql)
	if err != nil {
		log.Printf("Error getting DNS Records: %v", err)
		http.Error(out, "Unable to retrieve DNS Records", http.StatusInternalServerError)
	}
	contents := ""
	for r := range records {
		sections := strings.Split(records[r].IPv4, ".")

		for s := range sections {
			var ipnum int
			ipnum, _ = strconv.Atoi(sections[s])
			records[r].SortOrder += fmt.Sprintf("%03d", ipnum)
		}
	}

	sort.SliceStable(records, func(i, j int) bool {
		return records[i].SortOrder < records[j].SortOrder
	})

	contents += FileHeader
	for r := range records {
		contents += fmt.Sprintf("address=/%s/%s # %s\n", records[r].Hostname, records[r].IPv4, records[r].Label)

	}
	err = files.WriteOut("../config/02.dns.conf", contents)
	if err != nil {
		http.Error(out, "Unable to write DNS config file", http.StatusInternalServerError)
	}
}

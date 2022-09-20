package dnsmasq

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"facemasq/lib/db"
	"facemasq/lib/files"
)

type DHCP struct {
	MAC         string `db:"MAC"`
	IPv4        string `db:"IPv4"`
	Label       string `db:"Label"`
	MachineName string `db:"MachineName"`
	SortOrder   string
}

func WriteDHCPConfig(out http.ResponseWriter, in *http.Request) {
	var records []DHCP
	var exportDir string
	sql := `SELECT interfaces.mac, addresses.ipv4, devices.label, devices.machine_name FROM addresses JOIN interfaces ON interfaces.id = addresses.interface_id JOIN devices ON devices.id = interfaces.device_id WHERE addresses.is_reserved = 1;`
	err := db.Conn.NewRaw(sql).Scan(db.Context, &records)
	if err != nil {
		log.Printf("Error getting DHCP Records: %v", err)
		http.Error(out, "Unable to retrieve DHCP Records", http.StatusInternalServerError)
		return
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
		contents += fmt.Sprintf("dhcp-host=%s,%s # %s (%s)\n", records[r].MAC, records[r].IPv4, records[r].MachineName, records[r].Label)
	}
	exportDir, err = files.GetDir("export")
	if err != nil {
		http.Error(out, "Unable to determine where to export the DHCP config file", http.StatusInternalServerError)
	}

	err = files.WriteOut(fmt.Sprintf("%[2]s%[1]c%[3]s", os.PathSeparator, exportDir, DHCPFilename), contents)
	if err != nil {
		http.Error(out, "Unable to write DHCP config file", http.StatusInternalServerError)
	}
}

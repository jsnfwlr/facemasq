package handlers

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"facemasq/lib/db"
	"facemasq/lib/files"

	"github.com/uptrace/bunrouter"
)

type DHCP struct {
	MAC         string `bun:"mac"`
	IPv4        string `bun:"ipv4"`
	Label       string `bun:"label"`
	MachineName string `bun:"machine_name"`
	SortOrder   string
}

func WriteDHCPConfig(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var records []DHCP
	var exportDir string
	sql := `SELECT interfaces.mac, addresses.ipv4, devices.label, devices.machine_name FROM addresses JOIN interfaces ON interfaces.id = addresses.interface_id JOIN devices ON devices.id = interfaces.device_id WHERE addresses.is_reserved = 1;`
	err = db.Conn.NewRaw(sql).Scan(db.Context, &records)
	if err != nil {
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
		return
	}

	err = files.WriteOut(fmt.Sprintf("%[2]s%[1]c%[3]s", os.PathSeparator, exportDir, DHCPFilename), contents)
	if err != nil {
		return
	}
	return
}

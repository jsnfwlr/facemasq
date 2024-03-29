package handlers

import (
	"encoding/json"
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

type DNS struct {
	Hostname  string `bun:"hostname"`
	IPv4      string `bun:"ipv4"`
	Label     string `bun:"label"`
	SortOrder string
}

func WriteDNSConfig(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var exportDir string
	var records []DNS
	var formatHostnames string
	var suffixes []string

	err = db.Conn.NewRaw(`SELECT hostname, ipv4, devices.label FROM hostnames JOIN addresses ON hostnames.address_id = addresses.id JOIN interfaces ON addresses.interface_id = interfaces.id JOIN devices ON devices.id = interfaces.device_id WHERE is_dns = 1 ORDER BY hostname;`).Scan(db.Context, &records)
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

	err = db.Conn.NewRaw(`SELECT value FROM meta WHERE user_id IS NULL AND name = 'formatHostnames';`).Scan(db.Context, &formatHostnames)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(formatHostnames), &suffixes)
	if err != nil {
		return
	}

	contents += FileHeader
	for r := range records {
		for s := range suffixes {
			contents += fmt.Sprintf("address=/%s/%s # %s\n", fmt.Sprintf(suffixes[s], records[r].Hostname), records[r].IPv4, records[r].Label)
		}
	}

	exportDir, err = files.GetDir("export")
	if err != nil {
		return
	}

	err = files.WriteOut(fmt.Sprintf("%[2]s%[1]c%[3]s", os.PathSeparator, exportDir, DNSFilename), contents)
	if err != nil {
		return
	}
	return
}

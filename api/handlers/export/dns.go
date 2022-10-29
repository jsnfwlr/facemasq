package export

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
	"facemasq/lib/logging"
)

type DNS struct {
	Hostname  string `bun:"hostname"`
	IPv4      string `bun:"ipv4"`
	Label     string `bun:"label"`
	SortOrder string
}

func WriteDNSConfig(out http.ResponseWriter, in *http.Request) {
	var exportDir string
	var records []DNS
	var formatHostnames string
	var suffixes []string

	err := db.Conn.NewRaw(`SELECT hostname, ipv4, devices.label FROM hostnames JOIN addresses ON hostnames.address_id = addresses.id JOIN interfaces ON addresses.interface_id = interfaces.id JOIN devices ON devices.id = interfaces.device_id WHERE is_dns = 1 ORDER BY hostname;`).Scan(db.Context, &records)
	if err != nil {
		logging.Errorf("Error getting Hostname Records: %v", err)
		http.Error(out, "Unable to retrieve Hostname Records", http.StatusInternalServerError)
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
		logging.Errorf("Error getting Hostname suffixes: %v", err)
		http.Error(out, "Unable to retrieve Hostname suffixes", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal([]byte(formatHostnames), &suffixes)
	if err != nil {
		logging.Errorf("Error parsing Hostname suffixes: %v", err)
		http.Error(out, "Unable to parse Hostname suffixes", http.StatusInternalServerError)
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
		http.Error(out, "Unable to determine where to export the DNS config file", http.StatusInternalServerError)
	}

	err = files.WriteOut(fmt.Sprintf("%[2]s%[1]c%[3]s", os.PathSeparator, exportDir, DNSFilename), contents)
	if err != nil {
		http.Error(out, "Unable to write DNS config file", http.StatusInternalServerError)
	}
}

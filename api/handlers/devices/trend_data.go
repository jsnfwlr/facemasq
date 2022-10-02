package devices

import (
	"fmt"
	"net/http"
	"time"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/logging"

	"github.com/volatiletech/null"
)

type TrendWindow struct {
	Label   string
	from    time.Time
	to      time.Time
	Current int
	Compare int
	Tooltip null.String
}

type Concurrency struct {
	value int
	time  time.Time
}

func GetTrendData(out http.ResponseWriter, in *http.Request) {
	var firstSeen time.Time
	var concurrency Concurrency
	err := db.Conn.NewRaw("SELECT first_seen FROM devices ORDER BY first_seen ASC LIMIT 1;").Scan(db.Context, &firstSeen)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			logging.Errorln(err.Error())
			http.Error(out, "Error getting first_seen for oldest device", http.StatusInternalServerError)
			return
		}
	}
	durations := []TrendWindow{
		{
			Label:   "Historic",
			Current: 0,
			Compare: 0,
			Tooltip: null.StringFrom(fmt.Sprintf("Total number of unique devices since %s", firstSeen.Format("2006-01-02 15:04:05"))),
		},
		{
			Label:   "Concurrent",
			Current: 0,
			Compare: 0,
		},
		{
			Label:   "30 days",
			from:    time.Now().Add(time.Duration((60 * 24 * -1)) * time.Hour),
			to:      time.Now().Add(time.Duration((30 * 24 * -1)) * time.Hour),
			Current: 0,
			Compare: 0,
		},
		{
			Label:   "7 days",
			from:    time.Now().Add(time.Duration((14 * 24 * -1)) * time.Hour),
			to:      time.Now().Add(time.Duration((7 * 24 * -1)) * time.Hour),
			Current: 0,
			Compare: 0,
		},
		{
			Label:   "24 hours",
			from:    time.Now().Add(time.Duration(-48) * time.Hour),
			to:      time.Now().Add(time.Duration(-24) * time.Hour),
			Current: 0,
			Compare: 0,
		},
		{
			Label:   "30 minutes",
			from:    time.Now().Add(time.Duration(-60) * time.Minute),
			to:      time.Now().Add(time.Duration(-30) * time.Minute),
			Current: 0,
			Compare: 0,
		},
	}
	for d := range durations {
		switch durations[d].Label {
		case "Historic":
			err = db.Conn.NewRaw(`SELECT Count(active) as active FROM (SELECT DISTINCT devices.id as active, 1 as merge FROM devices JOIN interfaces ON interfaces.device_id = devices.id JOIN addresses ON addresses.interface_id = interfaces.id) as period GROUP BY merge;`).Scan(db.Context, &durations[d].Compare)
		case "Concurrent":
			err = db.Conn.NewRaw(`SELECT COUNT(address_id) as active FROM histories GROUP BY scan_id ORDER BY COUNT(address_id) DESC LIMIT 1 OFFSET 1;`).Scan(db.Context, &durations[d].Compare)
		default:
			err = db.Conn.NewRaw(`SELECT Count(active) as active, '' as Tooltip FROM (SELECT DISTINCT address_id as active, 1 as merge FROM histories JOIN scans ON histories.scan_id = scans.id WHERE scans.time > ? AND scans.time < ?) as period GROUP BY merge;`, durations[d].from.Format("2006-01-02 15:04"), durations[d].to.Format("2006-01-02 15:04")).Scan(db.Context, &durations[d].Compare)
		}

		if err != nil {
			if err.Error() != "sql: no rows in result set" {
				logging.Errorln(err.Error())
				http.Error(out, "Unable to retrieve comparative trend data for "+durations[d].Label, http.StatusInternalServerError)
				return
			}
			durations[d].Compare = 0
		}

		switch durations[d].Label {
		case "Historic":
			durations[d].Current = durations[d].Compare
		case "Concurrent":
			err = db.Conn.NewRaw(`SELECT active, scans.time FROM (SELECT COUNT(address_id) AS active, scan_id FROM histories GROUP BY scan_id ORDER BY COUNT(address_id) DESC LIMIT 1 OFFSET 0) AS peak JOIN scans ON scans.id = peak.scan_id;`).Scan(db.Context, &concurrency)
			durations[d].Current = concurrency.value
			durations[d].Tooltip = null.StringFrom(fmt.Sprintf("Peak number of unique devices online at one time (%s)", concurrency.time.Format("2006-01-02 15:04:05")))
		default:
			err = db.Conn.NewRaw(`SELECT Count(active) as active FROM (SELECT DISTINCT address_id as active, 1 as merge FROM histories JOIN scans ON histories.scan_id = scans.id WHERE scans.time > ? AND scans.time < ?) as period GROUP BY merge;`, durations[d].to.Format("2006-01-02 15:04"), time.Now().Format("2006-01-02 15:04")).Scan(db.Context, &durations[d].Current)
		}

		if err != nil {
			if err.Error() != "sql: no rows in result set" {
				logging.Errorln(err.Error())
				http.Error(out, "Unable to retrieve current trend data for "+durations[d].Label, http.StatusInternalServerError)
				return
			}
			durations[d].Current = 0
		}
	}

	formats.WriteJSONResponse(durations, out, in)
}

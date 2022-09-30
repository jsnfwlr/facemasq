package devices

import (
	"net/http"
	"time"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
)

type TrendWindow struct {
	Label   string
	from    time.Time
	to      time.Time
	Current int
	Compare int
}

func GetTrendData(out http.ResponseWriter, in *http.Request) {
	durations := []TrendWindow{
		{
			Label:   "Historic",
			Current: 0,
			Compare: 0,
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
	var sql string
	var err error
	for d := range durations {
		switch durations[d].Label {
		case "Historic":
			sql = `SELECT Count(active) as active FROM (SELECT DISTINCT devices.id as active, 1 as merge FROM devices JOIN interfaces ON interfaces.device_id = devices.id JOIN addresses ON addresses.interface_id = interfaces.id) as period GROUP BY merge;`
			err = db.Conn.NewRaw(sql).Scan(db.Context, &durations[d].Compare)
		case "Concurrent":
			sql = `SELECT COUNT(address_id) as active FROM histories GROUP BY scan_id ORDER BY COUNT(address_id) DESC LIMIT 1 OFFSET 1;`
			err = db.Conn.NewRaw(sql).Scan(db.Context, &durations[d].Compare)
		default:
			sql = `SELECT Count(active) as active FROM (SELECT DISTINCT address_id as active, 1 as merge FROM histories JOIN scans ON histories.scan_id = scans.id WHERE scans.time > ? AND scans.time < ?) as period GROUP BY merge;`
			err = db.Conn.NewRaw(sql, durations[d].from.Format("2006-01-02 15:04"), durations[d].to.Format("2006-01-02 15:04")).Scan(db.Context, &durations[d].Compare)
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
			sql = `SELECT COUNT(address_id) as active FROM histories GROUP BY scan_id ORDER BY COUNT(address_id) DESC LIMIT 1;`
			err = db.Conn.NewRaw(sql).Scan(db.Context, &durations[d].Current)
		default:
			sql = `SELECT Count(active) as active FROM (SELECT DISTINCT address_id as active, 1 as merge FROM histories JOIN scans ON histories.scan_id = scans.id WHERE scans.time > ? AND scans.time < ?) as period GROUP BY merge;`
			err = db.Conn.NewRaw(sql, durations[d].to.Format("2006-01-02 15:04"), time.Now().Format("2006-01-02 15:04")).Scan(db.Context, &durations[d].Current)
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

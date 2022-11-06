package devices

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
)

type DevicesOverTime struct {
	Time      time.Time `db:"time"`
	Addresses int       `db:"addresses"`
}

var frequency time.Duration

func init() {
	var err error
	frequency, err = time.ParseDuration(os.Getenv("CHART_FREQUENCY"))
	if err != nil {
		frequency = time.Duration(60) * time.Second
	}
}

func GetDashboardChartData(out http.ResponseWriter, in *http.Request) {
	var series map[string][]DevicesOverTime
	var err error

	series = make(map[string][]DevicesOverTime)

	series["full"], err = getAddressCountPerScan(time.Duration(-24) * time.Hour)
	if err != nil {
		logging.Error("error getting intial chart data: %v", err)
		http.Error(out, "Unable to retrieve inital chart data", http.StatusInternalServerError)
	}

	series["averaged"], err = getAverageAddressCountPerPeriod(time.Duration(-24)*time.Hour, time.Duration(5)*frequency)
	if err != nil {
		logging.Error("error getting averaged chart data: %v", err)
		http.Error(out, "Unable to retrieve averaged chart data", http.StatusInternalServerError)
	}

	formats.WriteJSONResponse(series, out, in)
}

func getAddressCountPerScan(overallTimeSpan time.Duration) (data []DevicesOverTime, err error) {
	sql := `SELECT scans.time, Count(histories.address_id) as addresses FROM scans JOIN histories ON histories.scan_id = scans.id WHERE time > ? GROUP BY time ORDER BY time ASC;`
	err = db.Conn.NewRaw(sql, time.Now().Add(overallTimeSpan).Format("2006-01-02 15:04")).Scan(db.Context, &data)

	return
}

func getAverageAddressCountPerPeriod(overallTimeSpan, averageIntervalSize time.Duration) (data []DevicesOverTime, err error) {
	var averageIntervals int
	var rawData []DevicesOverTime
	sql := `SELECT scans.time, Count(histories.address_id) as addresses FROM scans JOIN histories ON histories.scan_id = scans.id WHERE time > ? GROUP BY time ORDER BY time ASC;`
	err = db.Conn.NewRaw(sql, time.Now().Add(overallTimeSpan)).Scan(db.Context, &rawData)
	if err != nil {
		return
	}

	averageIntervals, err = strconv.Atoi(strings.Replace(((overallTimeSpan * -1) / averageIntervalSize).String(), "ns", "", -1))
	if err != nil {
		return
	}

	for i := 1; i <= averageIntervals; i++ {
		timeMin := time.Now().Add(overallTimeSpan)
		if i != 1 {
			timeMin = timeMin.Add(time.Duration((i - 1)) * averageIntervalSize)
		}
		timeMax := time.Now().Add(overallTimeSpan).Add(time.Duration(i) * averageIntervalSize)

		tally := []int{}
		first := time.Now()
		for j := range rawData {
			if (rawData[j].Time.Format("2006-01-02 15:04") == timeMin.Format("2006-01-02 15:04") || rawData[j].Time.After(timeMin)) && rawData[j].Time.Before(timeMax) {
				if len(tally) == 0 {
					first = rawData[j].Time
				}
				tally = append(tally, rawData[j].Addresses)
			}
			if rawData[j].Time.After(timeMax) {
				break
			}
		}
		sum := 0
		for j := range tally {
			sum += tally[j]
		}
		if len(tally) > 0 {
			avg := sum / len(tally)

			data = append(data, DevicesOverTime{Time: first, Addresses: avg})
		}
	}
	return
}

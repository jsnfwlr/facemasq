package devices

import (
	"fmt"
	"net/http"
	"time"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/models"
)

func SaveDevice(out http.ResponseWriter, in *http.Request) {
	var input models.Device
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		logging.Errorf("Unable to parse Device: %v", err)
		http.Error(out, "Unable to parse Device", http.StatusInternalServerError)
		return
	}
	if input.FirstSeen.Format("2006-01-02") == "0001-01-01" {
		input.FirstSeen = time.Now()
	}
	logging.Printf(2, "%v", input)
	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where("id = ?", input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		logging.Errorf("Unable to save Device: %v", err)
		http.Error(out, "Unable to save Device", http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteDevice(out http.ResponseWriter, in *http.Request) {
	var input models.Device
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		logging.Errorf("Unable to parse Device: %v", err)
		http.Error(out, "Unable to parse Device", http.StatusInternalServerError)
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	} else {
		err = fmt.Errorf("input (%v) is not a valid device record", input)
	}
	if err != nil {
		logging.Errorf("Unable to delete Device: %v", err)
		http.Error(out, "Unable to delete Device", http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

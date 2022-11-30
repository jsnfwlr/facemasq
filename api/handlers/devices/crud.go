package devices

import (
	"net/http"
	"time"

	"facemasq/lib/db"
	"facemasq/lib/devices"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/models"
)

func Save(out http.ResponseWriter, in *http.Request) {
	var input models.Device
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		logging.Error("Unable to parse Device: %v", err)
		http.Error(out, "Unable to parse Device", http.StatusInternalServerError)
		return
	}
	if input.FirstSeen.Format("2006-01-02") == "0001-01-01" {
		input.FirstSeen = time.Now()
	}
	logging.Debug("%v", input)
	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where("id = ?", input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		logging.Error("Unable to save Device: %v", err)
		http.Error(out, "Unable to save Device", http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func Delete(out http.ResponseWriter, in *http.Request) {
	var input models.Device
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		logging.Error("Unable to parse Device: %v", err)
		http.Error(out, "Unable to parse Device", http.StatusInternalServerError)
		return
	}

	err = devices.DeleteDevice(input.ID)
	if err != nil {
		logging.Error("Unable to delete Device: %v", err)
		http.Error(out, "Unable to delete Device", http.StatusInternalServerError)
		return

	}

	formats.WriteJSONResponse(input, out, in)
}

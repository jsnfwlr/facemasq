package devices

import (
	"net/http"

	"facemasq/lib/db"
	"facemasq/lib/devices"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/models"
)

func SaveInterface(out http.ResponseWriter, in *http.Request) {
	var input models.Interface
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		logging.Error("Unable to parse Inteface: %v", err)
		http.Error(out, "Unable to parse Inteface", http.StatusInternalServerError)
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where("id = ?", input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		logging.Error("Unable to save Inteface: %v", err)
		http.Error(out, "Unable to save Inteface", http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteInterface(out http.ResponseWriter, in *http.Request) {
	var input models.Device
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		logging.Error("Unable to parse Device: %v", err)
		http.Error(out, "Unable to parse Device", http.StatusInternalServerError)
		return
	}

	err = devices.DeleteInterface(input.ID)
	if err != nil {
		logging.Error("Unable to delete Device: %v", err)
		http.Error(out, "Unable to delete Device", http.StatusInternalServerError)
		return

	}

	formats.WriteJSONResponse(input, out, in)
}

package devices

import (
	"fmt"
	"net/http"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/models"
)

func SaveInterface(out http.ResponseWriter, in *http.Request) {
	var input models.Interface
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		logging.Errorf("Unable to parse Inteface: %v", err)
		http.Error(out, "Unable to parse Inteface", http.StatusInternalServerError)
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where("id = ?", input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		logging.Errorf("Unable to save Inteface: %v", err)
		http.Error(out, "Unable to save Inteface", http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteInterface(out http.ResponseWriter, in *http.Request) {
	var input models.Interface
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		logging.Errorf("Unable to parse Interface: %v", err)
		http.Error(out, "Unable to parse Interface", http.StatusInternalServerError)
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	} else {
		err = fmt.Errorf("input (%v) is not a valid interface record", input)
	}
	if err != nil {
		logging.Errorf("Unable to delete Interface: %v", err)
		http.Error(out, "Unable to delete Interface", http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

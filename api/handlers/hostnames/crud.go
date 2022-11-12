package hostnames

import (
	"fmt"
	"net/http"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/models"
)

func Save(out http.ResponseWriter, in *http.Request) {
	var input models.Hostname
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		logging.Error("Unable to parse Hostname: %v", err)
		http.Error(out, "Unable to parse Hostname", http.StatusInternalServerError)
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		logging.Error("Unable to save Hostname: %v", err)
		http.Error(out, "Unable to save Hostname", http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func Delete(out http.ResponseWriter, in *http.Request) {
	var input models.Hostname
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		logging.Error("Unable to parse Hostname: %v", err)
		http.Error(out, "Unable to parse Hostname", http.StatusInternalServerError)
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	} else {
		err = fmt.Errorf("input (%v) is not a valid hostname record", input)
	}
	if err != nil {
		logging.Error("Unable to delete Hostname: %v", err)
		http.Error(out, "Unable to delete Hostname", http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

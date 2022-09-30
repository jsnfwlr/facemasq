package devices

import (
	"fmt"
	"net/http"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/models"
)

func SaveAddress(out http.ResponseWriter, in *http.Request) {
	var input models.Address
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		logging.Errorf("Unable to parse Address: %v", err)
		http.Error(out, "Unable to parse Address", http.StatusInternalServerError)
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		logging.Errorf("Unable to save Address: %v", err)
		http.Error(out, "Unable to save Address", http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteAddress(out http.ResponseWriter, in *http.Request) {
	var input models.Address
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		logging.Errorf("Unable to parse Address: %v", err)
		http.Error(out, "Unable to parse Address", http.StatusInternalServerError)
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	} else {
		err = fmt.Errorf("input (%v) is not a valid address record", input)
	}
	if err != nil {
		logging.Errorf("Unable to delete Address: %v", err)
		http.Error(out, "Unable to delete Address", http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

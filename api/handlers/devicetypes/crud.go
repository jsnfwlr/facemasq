package devicetypes

import (
	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/lib/translate"
	"facemasq/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Save(out http.ResponseWriter, in *http.Request) {
	var input models.DeviceType
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		translation := translate.Message("SaveDeviceTypeError", "Unable to save DeviceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		translation := translate.Message("SaveDeviceTypeError", "Unable to save DeviceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func Delete(out http.ResponseWriter, in *http.Request) {
	var input models.DeviceType
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		translation := translate.Message("DeleteDeviceType", "Unable to delete DeviceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		translation := translate.Message("DeleteDeviceType", "Unable to delete DeviceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}
	formats.WriteJSONResponse(input, out, in)
}
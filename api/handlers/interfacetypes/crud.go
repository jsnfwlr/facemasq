package interfacetypes

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

func SaveInterfaceType(out http.ResponseWriter, in *http.Request) {
	var input models.InterfaceType
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		translation := translate.Message("SaveInterfaceTypeError", "Unable to save InterfaceType")
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
		translation := translate.Message("SaveInterfaceTypeError", "Unable to save InterfaceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteInterfaceType(out http.ResponseWriter, in *http.Request) {
	var input models.InterfaceType
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		translation := translate.Message("DeleteInterfaceType", "Unable to delete InterfaceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		translation := translate.Message("DeleteInterfaceType", "Unable to delete InterfaceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}
	formats.WriteJSONResponse(input, out, in)
}

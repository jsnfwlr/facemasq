package interfacetypes

import (
	"errors"
	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/translate"
	"facemasq/models"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func SaveInterfaceType(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var input models.InterfaceType
	err = formats.ReadJSONBody(in, &input)
	if err != nil {
		err = errors.New(translate.Message("SaveInterfaceTypeError", "Unable to save InterfaceType"))
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		err = errors.New(translate.Message("SaveInterfaceTypeError", "Unable to save InterfaceType"))
		return

	}
	formats.WriteJSONResponse(input, out, in)
	return
}

func DeleteInterfaceType(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var input models.InterfaceType
	var id int64
	id, err = in.Params().Int64("ID")
	if err != nil {
		err = errors.New(translate.Message("DeleteInterfaceType", "Unable to delete InterfaceType"))
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		err = errors.New(translate.Message("DeleteInterfaceType", "Unable to delete InterfaceType"))
		return
	}
	formats.WriteJSONResponse(input, out, in)
	return
}

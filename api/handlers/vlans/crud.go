package vlans

import (
	"errors"
	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/translate"
	"facemasq/models"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func Save(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var input models.VLAN
	err = formats.ReadJSONBody(in, &input)
	if err != nil {
		err = errors.New(translate.Message("SaveVLANError", "Unable to save VLAN"))
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		err = errors.New(translate.Message("SaveVLANError", "Unable to save VLAN"))
		return

	}
	formats.WriteJSONResponse(input, out, in)
	return
}

func Delete(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var input models.VLAN
	var id int64
	id, err = in.Params().Int64("ID")
	if err != nil {
		err = errors.New(translate.Message("DeleteVLAN", "Unable to delete VLAN"))
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		err = errors.New(translate.Message("DeleteVLAN", "Unable to delete VLAN"))
		return
	}
	formats.WriteJSONResponse(input, out, in)
	return
}

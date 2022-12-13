package users

import (
	"errors"
	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/translate"
	"facemasq/models"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func SaveUser(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var input models.User
	err = formats.ReadJSONBody(in, &input)
	if err != nil {
		err = errors.New(translate.Message("SaveUserError", "Unable to save User"))
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		err = errors.New(translate.Message("SaveUserError", "Unable to save User"))
		return

	}
	formats.WriteJSONResponse(input, out, in)
	return
}

func DeleteUser(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var input models.User
	var id int64
	id, err = in.Params().Int64("id")

	if err != nil {
		err = errors.New(translate.Message("DeleteUser", "Unable to delete User"))
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		err = errors.New(translate.Message("DeleteUser", "Unable to delete User"))
		return
	}
	formats.WriteJSONResponse(input, out, in)
	return
}

package maintainers

import (
	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/translate"
	"facemasq/models"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uptrace/bunrouter"
)

func Save(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var input models.Maintainer
	err = formats.ReadJSONBody(in, &input)
	if err != nil {
		err = errors.New(translate.Message("SaveMaintainerError", "Unable to save Maintainer"))
		// logging.Error("%s: %v", translation, err)
		// http.Error(out, translation, http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		err = errors.New(translate.Message("SaveMaintainerError", "Unable to save Maintainer"))
		return

	}
	formats.WriteJSONResponse(input, out, in)
	return
}

func Delete(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var input models.Maintainer
	var id int64
	id, err = in.Params().Int64("ID")
	if err != nil {
		err = errors.New(translate.Message("DeleteMaintainer", "Unable to delete Maintainer"))
		return
	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		err = errors.New(translate.Message("DeleteMaintainer", "Unable to delete Maintainer"))
		return
	}
	formats.WriteJSONResponse(input, out, in)
	return
}

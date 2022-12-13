package hostnames

import (
	"fmt"
	"net/http"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/models"

	"github.com/uptrace/bunrouter"
)

func Save(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var input models.Hostname
	err = formats.ReadJSONBody(in, &input)
	if err != nil {
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		return

	}
	formats.WriteJSONResponse(input, out, in)
	return
}

func Delete(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var input models.Hostname
	err = formats.ReadJSONBody(in, &input)
	if err != nil {
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	} else {
		err = fmt.Errorf("input (%v) is not a valid hostname record", input)
	}
	if err != nil {
		return

	}
	formats.WriteJSONResponse(input, out, in)
	return
}

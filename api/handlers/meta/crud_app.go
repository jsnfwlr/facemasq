package meta

import (
	"net/http"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/models"

	"github.com/uptrace/bunrouter"
	"github.com/volatiletech/null"
)

func GetAppSettings(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var settings []models.Meta
	err = db.Conn.NewSelect().Model(&settings).Where(`user_id IS NULL`).Scan(db.Context)
	if err != nil {
		logging.Error("error getting settings: %v", err)
		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
	}
	formats.WriteJSONResponse(settings, out, in)
	return
}

func SaveAppSetting(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var input, check models.Meta
	err = formats.ReadJSONBody(in, &input)
	if err != nil {
		return
	}

	input.UserID = null.Int64{Int64: 0, Valid: false}

	err = db.Conn.NewSelect().Model(&check).Where(`user_id IS NULL AND name = ?`, input.Name).Scan(db.Context)
	newSetting := false
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return
		}
		err = nil
		newSetting = true
	}
	if newSetting {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	} else {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`user_id IS NULL AND name = ?`, input.Name).Exec(db.Context)
	}
	if err != nil {
		return
	}

	formats.WriteJSONResponse(input, out, in)
	return
}

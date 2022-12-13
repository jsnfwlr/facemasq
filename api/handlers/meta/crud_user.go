package meta

import (
	"net/http"

	"github.com/uptrace/bunrouter"
	"github.com/volatiletech/null"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/models"
)

func GetUserSettings(out http.ResponseWriter, in bunrouter.Request) (err error) {
	userID := in.Params().ByName("userID")
	var settings []models.Meta
	sql := `SELECT name, value FROM meta WHERE user_id = ?`
	err = db.Conn.NewRaw(sql, userID).Scan(db.Context, &settings)
	if err != nil {
		return
	}
	formats.WriteJSONResponse(settings, out, in)
	return
}

func SaveUserSetting(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var input, check models.Meta
	var userID int64
	userID, err = in.Params().Int64("userID")
	if err != nil {
		return
	}

	err = formats.ReadJSONBody(in, &input)
	if err != nil {
		return
	}

	input.UserID = null.Int64{Int64: userID, Valid: true}

	sql := `SELECT name, value, user_id FROM meta WHERE user_id = ? AND name = ?;`
	err = db.Conn.NewRaw(sql, input.UserID.Int64, input.Name).Scan(db.Context, &check)

	newSetting := false
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			logging.Error("Cant find %s,%d", input.Name, input.UserID.Int64)
			return
		}
		newSetting = true
		err = nil
	}
	if newSetting {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	} else {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`name = ? and user_id = ?`, input.Name, input.UserID.Int64).Exec(db.Context)
	}
	if err != nil {
		return
	}

	formats.WriteJSONResponse(input, out, in)
	return
}

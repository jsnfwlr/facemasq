package meta

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/volatiletech/null"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/models"
)

func GetUserSettings(out http.ResponseWriter, in *http.Request) {
	userID := mux.Vars(in)["userID"]
	var settings []models.Meta
	sql := `SELECT name, value FROM meta WHERE user_id = ?`
	err := db.Conn.NewRaw(sql, userID).Scan(db.Context, &settings)
	if err != nil {
		logging.Error("error getting settings: %v", err)
		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
	}
	formats.WriteJSONResponse(settings, out, in)
}

func SaveUserSetting(out http.ResponseWriter, in *http.Request) {
	var input, check models.Meta

	userID, err := strconv.ParseInt(mux.Vars(in)["userID"], 10, 64)
	if err != nil {
		logging.Error("Unable to parse user_id: %v", err)
		http.Error(out, "Unable to parse UserID", http.StatusBadRequest)
		return
	}

	err = formats.ReadJSONBody(in, &input)
	if err != nil {
		logging.Error("Unable to parse Setting: %v", err)
		http.Error(out, "Unable to parse Setting", http.StatusInternalServerError)
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
	}
	if newSetting {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	} else {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`name = ? and user_id = ?`, input.Name, input.UserID.Int64).Exec(db.Context)
	}
	if err != nil {
		logging.Error("error saving setting: %v", err)
		http.Error(out, "Unable to store data", http.StatusInternalServerError)
	}

	formats.WriteJSONResponse(input, out, in)
}

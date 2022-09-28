package settings

import (
	"log"
	"net/http"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/models"

	"github.com/volatiletech/null"
)

func GetAppSettings(out http.ResponseWriter, in *http.Request) {
	var settings []models.Meta
	err := db.Conn.NewSelect().Model(&settings).Where(`user_id IS NULL`).Scan(db.Context)
	if err != nil {
		log.Printf("error getting settings: %v", err)
		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
	}
	formats.WriteJSONResponse(settings, out, in)
}

func SaveAppSetting(out http.ResponseWriter, in *http.Request) {
	var input, check models.Meta
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		log.Printf("Unable to parse Setting: %v", err)
		http.Error(out, "Unable to parse Setting", http.StatusInternalServerError)
		return
	}

	input.UserID = null.Int64{Int64: 0, Valid: false}

	err = db.Conn.NewSelect().Model(&check).Where(`user_id IS NULL AND name = ?`, input.Name).Scan(db.Context)
	newSetting := false
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return
		}
		newSetting = true
	}
	if newSetting {
		_, err = db.Conn.NewUpdate().Model(&input).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("error saving setting: %v", err)
		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
	}

	formats.WriteJSONResponse(input, out, in)
}

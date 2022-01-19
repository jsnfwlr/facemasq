package settings

import (
	"log"
	"net/http"

	"github.com/volatiletech/null"

	"github.com/jsnfwlr/facemasq/api/lib/db"
	"github.com/jsnfwlr/facemasq/api/lib/formats"
	"github.com/jsnfwlr/facemasq/api/models/Meta"
)

func GetAppSettings(out http.ResponseWriter, in *http.Request) {
	var settings []Meta.Model
	sql := `SELECT Name, Value, UserID  FROM Meta WHERE UserID = NULL`
	err := db.Conn.Select(&settings, sql)
	if err != nil {
		log.Printf("error getting settings: %v", err)
		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
	}
	formats.PublishJSON(settings, out, in)
}

func SaveAppSetting(out http.ResponseWriter, in *http.Request) {
	var input Meta.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to parse Setting: %v", err)
		http.Error(out, "Unable to parse Setting", http.StatusInternalServerError)
		return
	}

	input.UserID = null.NewInt64(0, false)

	err = input.Save()
	if err != nil {
		log.Printf("error saving setting: %v", err)
		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
	}

	formats.PublishJSON(input, out, in)
}

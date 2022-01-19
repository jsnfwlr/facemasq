package settings

import (
	"net/http"
	"log"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/volatiletech/null"

	"github.com/jsnfwlr/facemasq/api/lib/db"
	"github.com/jsnfwlr/facemasq/api/lib/formats"
	"github.com/jsnfwlr/facemasq/api/models/Meta"
)

func GetUserSettings(out http.ResponseWriter, in *http.Request) {
	userID := mux.Vars(in)["userID"]
	var settings []Meta.Model
	sql := `SELECT Name, Value FROM Meta WHERE UserID = ?`
	err := db.Conn.Select(&settings, sql, userID)
	if err != nil {
		log.Printf("error getting settings: %v", err)
		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
	}
	formats.PublishJSON(settings, out, in)
}

func SaveUserSetting(out http.ResponseWriter, in *http.Request) {
	var input Meta.Model

	userID, err := strconv.ParseInt(mux.Vars(in)["userID"], 10, 64)
	if err != nil {
	        log.Printf("Unable to parse UserID: %v", err)
                http.Error(out, "Unable to parse UserID", http.StatusBadRequest)
                return
	}

        err = formats.ReadJSON(in, &input)
        if err != nil {
                log.Printf("Unable to parse Setting: %v", err)
                http.Error(out, "Unable to parse Setting", http.StatusInternalServerError)
                return
        }

	input.UserID = null.Int64From(userID)

	err = input.Save()
	if err != nil {
		log.Printf("error saving setting: %v", err)
		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
	}
	

	formats.PublishJSON(input, out, in)
}


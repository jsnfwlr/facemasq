package devices

import (
	"net/http"
	"time"

	"facemasq/lib/db"
	"facemasq/lib/devices"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/models"

	"github.com/uptrace/bunrouter"
)

func Save(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var input models.Device
	err = formats.ReadJSONBody(in, &input)
	if err != nil {
		return
	}
	if input.FirstSeen.Format("2006-01-02") == "0001-01-01" {
		input.FirstSeen = time.Now()
	}
	logging.Debug("%v", input)
	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where("id = ?", input.ID).Exec(db.Context)
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
	var input models.Device
	err = formats.ReadJSONBody(in, &input)
	if err != nil {
		return
	}

	err = devices.DeleteDevice(input.ID)
	if err != nil {
		return
	}

	formats.WriteJSONResponse(input, out, in)
	return
}

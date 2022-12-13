package params

import (
	"net/http"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/models"

	"github.com/uptrace/bunrouter"
)

type Params struct {
	Categories       []models.Category
	Statuses         []models.Status
	Locations        []models.Location
	Maintainers      []models.Maintainer
	Architectures    []models.Architecture
	OperatingSystems []models.OperatingSystem
	InterfaceTypes   []models.InterfaceType
	DeviceTypes      []models.DeviceType
	VLANs            []models.VLAN
	Users            []models.User
}

func GetAllParams(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var params Params

	err = db.Conn.NewSelect().Model(&params.Categories).Scan(db.Context)
	if err != nil {
		return
	}

	err = db.Conn.NewSelect().Model(&params.Statuses).Scan(db.Context)
	if err != nil {
		return
	}

	err = db.Conn.NewSelect().Model(&params.Locations).Scan(db.Context)
	if err != nil {
		return
	}

	err = db.Conn.NewSelect().Model(&params.Maintainers).Scan(db.Context)
	if err != nil {
		return
	}

	err = db.Conn.NewSelect().Model(&params.Architectures).Scan(db.Context)
	if err != nil {
		return
	}

	err = db.Conn.NewSelect().Model(&params.OperatingSystems).Scan(db.Context)
	if err != nil {
		return
	}

	err = db.Conn.NewSelect().Model(&params.InterfaceTypes).Scan(db.Context)
	if err != nil {
		return
	}

	err = db.Conn.NewSelect().Model(&params.DeviceTypes).Scan(db.Context)
	if err != nil {
		return
	}

	err = db.Conn.NewSelect().Model(&params.VLANs).Scan(db.Context)
	if err != nil {
		return
	}

	err = db.Conn.NewSelect().Model(&params.Users).Scan(db.Context)
	if err != nil {
		return
	}

	formats.WriteJSONResponse(params, out, in)
	return
}

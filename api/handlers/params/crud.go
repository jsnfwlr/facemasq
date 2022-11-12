package params

import (
	"net/http"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/translate"
	"facemasq/models"
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

func GetAllParams(out http.ResponseWriter, in *http.Request) {
	var params Params
	var err error
	err = db.Conn.NewSelect().Model(&params.Categories).Scan(db.Context)
	if err != nil {
		http.Error(out, translate.Message("RetrieveCategoryError", "Unable to retrieve Category data"), http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.Statuses).Scan(db.Context)
	if err != nil {
		http.Error(out, translate.Message("RetrieveStatusError", "Unable to retrieve Status data"), http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.Locations).Scan(db.Context)
	if err != nil {
		http.Error(out, translate.Message("RetrieveLocationError", "Unable to retrieve Location data"), http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.Maintainers).Scan(db.Context)
	if err != nil {
		http.Error(out, translate.Message("RetrieveMaintainerError", "Unable to retrieve Maintainer data"), http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.Architectures).Scan(db.Context)
	if err != nil {
		http.Error(out, translate.Message("RetrieveArchitectureError", "Unable to retrieve Architecture data"), http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.OperatingSystems).Scan(db.Context)
	if err != nil {
		http.Error(out, translate.Message("RetrieveOperatingSystemError", "Unable to retrieve OperatingSystem data"), http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.InterfaceTypes).Scan(db.Context)
	if err != nil {
		http.Error(out, translate.Message("RetrieveInterfaceTypeError", "Unable to retrieve InterfaceType data"), http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.DeviceTypes).Scan(db.Context)
	if err != nil {
		http.Error(out, translate.Message("RetrieveDeviceTypeError", "Unable to retrieve DeviceType data"), http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.VLANs).Scan(db.Context)
	if err != nil {
		http.Error(out, translate.Message("RetrieveVLANError", "Unable to retrieve VLAN data"), http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.Users).Scan(db.Context)
	if err != nil {
		http.Error(out, translate.Message("RetrieveUserError", "Unable to retrieve User data"), http.StatusInternalServerError)
	}

	formats.WriteJSONResponse(params, out, in)
}

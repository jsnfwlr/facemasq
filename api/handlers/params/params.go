package params

import (
	"net/http"
	"strconv"

	"facemasq/lib/db"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/lib/translate"
	"facemasq/models"

	"github.com/gorilla/mux"
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

func GetAll(out http.ResponseWriter, in *http.Request) {
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

func SaveCategory(out http.ResponseWriter, in *http.Request) {
	var input models.Category
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		translation := translate.Message("SaveCategoryError", "Unable to save Category")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		translation := translate.Message("SaveCategoryError", "Unable to save Category")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteCategory(out http.ResponseWriter, in *http.Request) {
	var input models.Category
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		translation := translate.Message("DeleteCategory", "Unable to delete Category")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		translation := translate.Message("DeleteCategory", "Unable to delete Category")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}
	formats.WriteJSONResponse(input, out, in)
}

func SaveStatus(out http.ResponseWriter, in *http.Request) {
	var input models.Status
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		translation := translate.Message("SaveStatusError", "Unable to save Status")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		translation := translate.Message("SaveStatusError", "Unable to save Status")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteStatus(out http.ResponseWriter, in *http.Request) {
	var input models.Status
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		translation := translate.Message("DeleteStatus", "Unable to delete Status")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		translation := translate.Message("DeleteStatus", "Unable to delete Status")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}
	formats.WriteJSONResponse(input, out, in)
}

func SaveLocation(out http.ResponseWriter, in *http.Request) {
	var input models.Location
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		translation := translate.Message("SaveLocationError", "Unable to save Location")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		translation := translate.Message("SaveLocationError", "Unable to save Location")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteLocation(out http.ResponseWriter, in *http.Request) {
	var input models.Location
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		translation := translate.Message("DeleteLocation", "Unable to delete Location")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		translation := translate.Message("DeleteLocation", "Unable to delete Location")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}
	formats.WriteJSONResponse(input, out, in)
}

func SaveMaintainer(out http.ResponseWriter, in *http.Request) {
	var input models.Maintainer
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		translation := translate.Message("SaveMaintainerError", "Unable to save Maintainer")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		translation := translate.Message("SaveMaintainerError", "Unable to save Maintainer")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteMaintainer(out http.ResponseWriter, in *http.Request) {
	var input models.Maintainer
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		translation := translate.Message("DeleteMaintainer", "Unable to delete Maintainer")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		translation := translate.Message("DeleteMaintainer", "Unable to delete Maintainer")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}
	formats.WriteJSONResponse(input, out, in)
}

func SaveArchitecture(out http.ResponseWriter, in *http.Request) {
	var input models.Architecture
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		translation := translate.Message("SaveArchitectureError", "Unable to save Architecture")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		translation := translate.Message("SaveArchitectureError", "Unable to save Architecture")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteArchitecture(out http.ResponseWriter, in *http.Request) {
	var input models.Architecture
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		translation := translate.Message("DeleteArchitecture", "Unable to delete Architecture")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		translation := translate.Message("DeleteArchitecture", "Unable to delete Architecture")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}
	formats.WriteJSONResponse(input, out, in)
}

func SaveOperatingSystem(out http.ResponseWriter, in *http.Request) {
	var input models.OperatingSystem
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		translation := translate.Message("SaveOperatingSystemError", "Unable to save OperatingSystem")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		translation := translate.Message("SaveOperatingSystemError", "Unable to save OperatingSystem")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteOperatingSystem(out http.ResponseWriter, in *http.Request) {
	var input models.OperatingSystem
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		translation := translate.Message("DeleteOperatingSystem", "Unable to delete OperatingSystem")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		translation := translate.Message("DeleteOperatingSystem", "Unable to delete OperatingSystem")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}
	formats.WriteJSONResponse(input, out, in)
}

func SaveInterfaceType(out http.ResponseWriter, in *http.Request) {
	var input models.InterfaceType
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		translation := translate.Message("SaveInterfaceTypeError", "Unable to save InterfaceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		translation := translate.Message("SaveInterfaceTypeError", "Unable to save InterfaceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteInterfaceType(out http.ResponseWriter, in *http.Request) {
	var input models.InterfaceType
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		translation := translate.Message("DeleteInterfaceType", "Unable to delete InterfaceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		translation := translate.Message("DeleteInterfaceType", "Unable to delete InterfaceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}
	formats.WriteJSONResponse(input, out, in)
}

func SaveDeviceType(out http.ResponseWriter, in *http.Request) {
	var input models.DeviceType
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		translation := translate.Message("SaveDeviceTypeError", "Unable to save DeviceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		translation := translate.Message("SaveDeviceTypeError", "Unable to save DeviceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteDeviceType(out http.ResponseWriter, in *http.Request) {
	var input models.DeviceType
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		translation := translate.Message("DeleteDeviceType", "Unable to delete DeviceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		translation := translate.Message("DeleteDeviceType", "Unable to delete DeviceType")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}
	formats.WriteJSONResponse(input, out, in)
}

func SaveVLAN(out http.ResponseWriter, in *http.Request) {
	var input models.VLAN
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		translation := translate.Message("SaveVLANError", "Unable to save VLAN")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		translation := translate.Message("SaveVLANError", "Unable to save VLAN")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteVLAN(out http.ResponseWriter, in *http.Request) {
	var input models.VLAN
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		translation := translate.Message("DeleteVLAN", "Unable to delete VLAN")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		translation := translate.Message("DeleteVLAN", "Unable to delete VLAN")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}
	formats.WriteJSONResponse(input, out, in)
}

func SaveUser(out http.ResponseWriter, in *http.Request) {
	var input models.User
	err := formats.ReadJSONBody(in, &input)
	if err != nil {
		translation := translate.Message("SaveUserError", "Unable to save User")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Where(`id = ?`, input.ID).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		translation := translate.Message("SaveUserError", "Unable to save User")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	formats.WriteJSONResponse(input, out, in)
}

func DeleteUser(out http.ResponseWriter, in *http.Request) {
	var input models.User
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		translation := translate.Message("DeleteUser", "Unable to delete User")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		translation := translate.Message("DeleteUser", "Unable to delete User")
		logging.Error("%s: %v", translation, err)
		http.Error(out, translation, http.StatusInternalServerError)
		return
	}
	formats.WriteJSONResponse(input, out, in)
}

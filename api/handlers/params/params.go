package params

import (
	"log"
	"net/http"
	"strconv"

	"facemasq/lib/db"
	"facemasq/lib/formats"
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
		http.Error(out, "Unable to Category data", http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.Statuses).Scan(db.Context)
	if err != nil {
		http.Error(out, "Unable to Status data", http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.Locations).Scan(db.Context)
	if err != nil {
		http.Error(out, "Unable to Location data", http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.Maintainers).Scan(db.Context)
	if err != nil {
		http.Error(out, "Unable to Maintainer data", http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.Architectures).Scan(db.Context)
	if err != nil {
		http.Error(out, "Unable to Architecture data", http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.OperatingSystems).Scan(db.Context)
	if err != nil {
		http.Error(out, "Unable to OperatingSystem data", http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.InterfaceTypes).Scan(db.Context)
	if err != nil {
		http.Error(out, "Unable to InterfaceType data", http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.DeviceTypes).Scan(db.Context)
	if err != nil {
		http.Error(out, "Unable to DeviceType data", http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.VLANs).Scan(db.Context)
	if err != nil {
		http.Error(out, "Unable to VLAN data", http.StatusInternalServerError)
	}

	err = db.Conn.NewSelect().Model(&params.Users).Scan(db.Context)
	if err != nil {
		http.Error(out, "Unable to User data", http.StatusInternalServerError)
	}

	formats.PublishJSON(params, out, in)
}

func SaveCategory(out http.ResponseWriter, in *http.Request) {
	var input models.Category
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save Category: %v", err)
		http.Error(out, "Unable to save Category", http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("Unable to save Category: %v", err)
		http.Error(out, "Unable to save Category", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}

func DeleteCategory(out http.ResponseWriter, in *http.Request) {
	var input models.Category
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete Category: %v", err)
		http.Error(out, "Unable to delete Category", http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		log.Printf("Unable to delete Category: %v", err)
		http.Error(out, "Unable to delete Category", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveStatus(out http.ResponseWriter, in *http.Request) {
	var input models.Status
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save Status: %v", err)
		http.Error(out, "Unable to save Status", http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("Unable to save Status: %v", err)
		http.Error(out, "Unable to save Status", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}

func DeleteStatus(out http.ResponseWriter, in *http.Request) {
	var input models.Status
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete Status: %v", err)
		http.Error(out, "Unable to delete Status", http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		log.Printf("Unable to delete Status: %v", err)
		http.Error(out, "Unable to delete Status", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveLocation(out http.ResponseWriter, in *http.Request) {
	var input models.Location
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save Location: %v", err)
		http.Error(out, "Unable to save Location", http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("Unable to save Location: %v", err)
		http.Error(out, "Unable to save Location", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}

func DeleteLocation(out http.ResponseWriter, in *http.Request) {
	var input models.Location
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete Location: %v", err)
		http.Error(out, "Unable to delete Location", http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		log.Printf("Unable to delete Location: %v", err)
		http.Error(out, "Unable to delete Location", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveMaintainer(out http.ResponseWriter, in *http.Request) {
	var input models.Maintainer
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save Maintainer: %v", err)
		http.Error(out, "Unable to save Maintainer", http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("Unable to save Maintainer: %v", err)
		http.Error(out, "Unable to save Maintainer", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}

func DeleteMaintainer(out http.ResponseWriter, in *http.Request) {
	var input models.Maintainer
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete Maintainer: %v", err)
		http.Error(out, "Unable to delete Maintainer", http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		log.Printf("Unable to delete Maintainer: %v", err)
		http.Error(out, "Unable to delete Maintainer", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveArchitecture(out http.ResponseWriter, in *http.Request) {
	var input models.Architecture
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save Architecture: %v", err)
		http.Error(out, "Unable to save Architecture", http.StatusInternalServerError)
		return
	}
	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("Unable to save Architecture: %v", err)
		http.Error(out, "Unable to save Architecture", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}

func DeleteArchitecture(out http.ResponseWriter, in *http.Request) {
	var input models.Architecture
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete Architecture: %v", err)
		http.Error(out, "Unable to delete Architecture", http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		log.Printf("Unable to delete Architecture: %v", err)
		http.Error(out, "Unable to delete Architecture", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveOperatingSystem(out http.ResponseWriter, in *http.Request) {
	var input models.OperatingSystem
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save OperatingSystem: %v", err)
		http.Error(out, "Unable to save OperatingSystem", http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("Unable to save OperatingSystem: %v", err)
		http.Error(out, "Unable to save OperatingSystem", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}

func DeleteOperatingSystem(out http.ResponseWriter, in *http.Request) {
	var input models.OperatingSystem
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete OperatingSystem: %v", err)
		http.Error(out, "Unable to delete OperatingSystem", http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		log.Printf("Unable to delete OperatingSystem: %v", err)
		http.Error(out, "Unable to delete OperatingSystem", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveInterfaceType(out http.ResponseWriter, in *http.Request) {
	var input models.InterfaceType
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save InterfaceType: %v", err)
		http.Error(out, "Unable to save InterfaceType", http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("Unable to save InterfaceType: %v", err)
		http.Error(out, "Unable to save InterfaceType", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}

func DeleteInterfaceType(out http.ResponseWriter, in *http.Request) {
	var input models.InterfaceType
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete InterfaceType: %v", err)
		http.Error(out, "Unable to delete InterfaceType", http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		log.Printf("Unable to delete InterfaceType: %v", err)
		http.Error(out, "Unable to delete InterfaceType", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveDeviceType(out http.ResponseWriter, in *http.Request) {
	var input models.DeviceType
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save DeviceType: %v", err)
		http.Error(out, "Unable to save DeviceType", http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("Unable to save DeviceType: %v", err)
		http.Error(out, "Unable to save DeviceType", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}

func DeleteDeviceType(out http.ResponseWriter, in *http.Request) {
	var input models.DeviceType
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete DeviceType: %v", err)
		http.Error(out, "Unable to delete DeviceType", http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		log.Printf("Unable to delete DeviceType: %v", err)
		http.Error(out, "Unable to delete DeviceType", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveVLAN(out http.ResponseWriter, in *http.Request) {
	var input models.VLAN
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save VLAN: %v", err)
		http.Error(out, "Unable to save VLAN", http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("Unable to save VLAN: %v", err)
		http.Error(out, "Unable to save VLAN", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}

func DeleteVLAN(out http.ResponseWriter, in *http.Request) {
	var input models.VLAN
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete VLAN: %v", err)
		http.Error(out, "Unable to delete VLAN", http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		log.Printf("Unable to delete VLAN: %v", err)
		http.Error(out, "Unable to delete VLAN", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveUser(out http.ResponseWriter, in *http.Request) {
	var input models.User
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save User: %v", err)
		http.Error(out, "Unable to save User", http.StatusInternalServerError)
		return
	}

	if input.ID > 0 {
		_, err = db.Conn.NewUpdate().Model(&input).Exec(db.Context)
	} else {
		_, err = db.Conn.NewInsert().Model(&input).Exec(db.Context)
	}
	if err != nil {
		log.Printf("Unable to save User: %v", err)
		http.Error(out, "Unable to save User", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}

func DeleteUser(out http.ResponseWriter, in *http.Request) {
	var input models.User
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete User: %v", err)
		http.Error(out, "Unable to delete User", http.StatusInternalServerError)
		return

	}
	input.ID = id

	_, err = db.Conn.NewDelete().Model(&input).WherePK().Exec(db.Context)
	if err != nil {
		log.Printf("Unable to delete User: %v", err)
		http.Error(out, "Unable to delete User", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

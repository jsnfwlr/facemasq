package params

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jsnfwlr/facemasq/api/lib/formats"
	"github.com/jsnfwlr/facemasq/api/models/Architectures"
	"github.com/jsnfwlr/facemasq/api/models/Categories"
	"github.com/jsnfwlr/facemasq/api/models/DeviceTypes"
	"github.com/jsnfwlr/facemasq/api/models/InterfaceTypes"
	"github.com/jsnfwlr/facemasq/api/models/Locations"
	"github.com/jsnfwlr/facemasq/api/models/Maintainers"
	"github.com/jsnfwlr/facemasq/api/models/OperatingSystems"
	"github.com/jsnfwlr/facemasq/api/models/Status"
	"github.com/jsnfwlr/facemasq/api/models/Users"
	"github.com/jsnfwlr/facemasq/api/models/VLANs"
)

type Params struct {
	Categories       []Categories.Model
	Statuses         []Status.Model
	Locations        []Locations.Model
	Maintainers      []Maintainers.Model
	Architectures    []Architectures.Model
	OperatingSystems []OperatingSystems.Model
	InterfaceTypes   []InterfaceTypes.Model
	DeviceTypes      []DeviceTypes.Model
	VLANs            []VLANs.Model
	Users            []Users.Model
}

func GetAll(out http.ResponseWriter, in *http.Request) {
	var params Params
	var err error
	params.Categories, err = Categories.Get()
	if err != nil {
		http.Error(out, "Unable to Category data", http.StatusInternalServerError)
	}

	params.Statuses, err = Status.Get()
	if err != nil {
		http.Error(out, "Unable to Status data", http.StatusInternalServerError)
	}

	params.Locations, err = Locations.Get()
	if err != nil {
		http.Error(out, "Unable to Location data", http.StatusInternalServerError)
	}

	params.Maintainers, err = Maintainers.Get()
	if err != nil {
		http.Error(out, "Unable to Maintainer data", http.StatusInternalServerError)
	}

	params.Architectures, err = Architectures.Get()
	if err != nil {
		http.Error(out, "Unable to Architecture data", http.StatusInternalServerError)
	}

	params.OperatingSystems, err = OperatingSystems.Get()
	if err != nil {
		http.Error(out, "Unable to OperatingSystem data", http.StatusInternalServerError)
	}

	params.InterfaceTypes, err = InterfaceTypes.Get()
	if err != nil {
		http.Error(out, "Unable to InterfaceType data", http.StatusInternalServerError)
	}

	params.DeviceTypes, err = DeviceTypes.Get()
	if err != nil {
		http.Error(out, "Unable to DeviceType data", http.StatusInternalServerError)
	}

	params.VLANs, err = VLANs.Get()
	if err != nil {
		http.Error(out, "Unable to VLAN data", http.StatusInternalServerError)
	}

	params.Users, err = Users.Get()
	if err != nil {
		http.Error(out, "Unable to User data", http.StatusInternalServerError)
	}

	formats.PublishJSON(params, out, in)
}

func SaveCategory(out http.ResponseWriter, in *http.Request) {
	var input Categories.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save Category: %v", err)
		http.Error(out, "Unable to save Category", http.StatusInternalServerError)
		return
	}
	err = input.Save()
	if err != nil {
		log.Printf("Unable to save Category: %v", err)
		http.Error(out, "Unable to save Category", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}
func DeleteCategory(out http.ResponseWriter, in *http.Request) {
	var input Categories.Model
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete Category: %v", err)
		http.Error(out, "Unable to delete Category", http.StatusInternalServerError)
		return

	}
	input.ID = id

	err = input.Delete()
	if err != nil {
		log.Printf("Unable to delete Category: %v", err)
		http.Error(out, "Unable to delete Category", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveStatus(out http.ResponseWriter, in *http.Request) {
	var input Status.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save Status: %v", err)
		http.Error(out, "Unable to save Status", http.StatusInternalServerError)
		return
	}
	err = input.Save()
	if err != nil {
		log.Printf("Unable to save Status: %v", err)
		http.Error(out, "Unable to save Status", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}
func DeleteStatus(out http.ResponseWriter, in *http.Request) {
	var input Status.Model
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete Status: %v", err)
		http.Error(out, "Unable to delete Status", http.StatusInternalServerError)
		return

	}
	input.ID = id

	err = input.Delete()
	if err != nil {
		log.Printf("Unable to delete Status: %v", err)
		http.Error(out, "Unable to delete Status", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveLocation(out http.ResponseWriter, in *http.Request) {
	var input Locations.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save Location: %v", err)
		http.Error(out, "Unable to save Location", http.StatusInternalServerError)
		return
	}
	err = input.Save()
	if err != nil {
		log.Printf("Unable to save Location: %v", err)
		http.Error(out, "Unable to save Location", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}
func DeleteLocation(out http.ResponseWriter, in *http.Request) {
	var input Locations.Model
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete Location: %v", err)
		http.Error(out, "Unable to delete Location", http.StatusInternalServerError)
		return

	}
	input.ID = id

	err = input.Delete()
	if err != nil {
		log.Printf("Unable to delete Location: %v", err)
		http.Error(out, "Unable to delete Location", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveMaintainer(out http.ResponseWriter, in *http.Request) {
	var input Maintainers.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save Maintainer: %v", err)
		http.Error(out, "Unable to save Maintainer", http.StatusInternalServerError)
		return
	}
	err = input.Save()
	if err != nil {
		log.Printf("Unable to save Maintainer: %v", err)
		http.Error(out, "Unable to save Maintainer", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}
func DeleteMaintainer(out http.ResponseWriter, in *http.Request) {
	var input Maintainers.Model
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete Maintainer: %v", err)
		http.Error(out, "Unable to delete Maintainer", http.StatusInternalServerError)
		return

	}
	input.ID = id

	err = input.Delete()
	if err != nil {
		log.Printf("Unable to delete Maintainer: %v", err)
		http.Error(out, "Unable to delete Maintainer", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveArchitecture(out http.ResponseWriter, in *http.Request) {
	var input Architectures.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save Architecture: %v", err)
		http.Error(out, "Unable to save Architecture", http.StatusInternalServerError)
		return
	}
	err = input.Save()
	if err != nil {
		log.Printf("Unable to save Architecture: %v", err)
		http.Error(out, "Unable to save Architecture", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}
func DeleteArchitecture(out http.ResponseWriter, in *http.Request) {
	var input Architectures.Model
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete Architecture: %v", err)
		http.Error(out, "Unable to delete Architecture", http.StatusInternalServerError)
		return

	}
	input.ID = id

	err = input.Delete()
	if err != nil {
		log.Printf("Unable to delete Architecture: %v", err)
		http.Error(out, "Unable to delete Architecture", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveOperatingSystem(out http.ResponseWriter, in *http.Request) {
	var input OperatingSystems.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save OperatingSystem: %v", err)
		http.Error(out, "Unable to save OperatingSystem", http.StatusInternalServerError)
		return
	}
	err = input.Save()
	if err != nil {
		log.Printf("Unable to save OperatingSystem: %v", err)
		http.Error(out, "Unable to save OperatingSystem", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}
func DeleteOperatingSystem(out http.ResponseWriter, in *http.Request) {
	var input OperatingSystems.Model
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete OperatingSystem: %v", err)
		http.Error(out, "Unable to delete OperatingSystem", http.StatusInternalServerError)
		return

	}
	input.ID = id

	err = input.Delete()
	if err != nil {
		log.Printf("Unable to delete OperatingSystem: %v", err)
		http.Error(out, "Unable to delete OperatingSystem", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveInterfaceType(out http.ResponseWriter, in *http.Request) {
	var input InterfaceTypes.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save InterfaceType: %v", err)
		http.Error(out, "Unable to save InterfaceType", http.StatusInternalServerError)
		return
	}
	err = input.Save()
	if err != nil {
		log.Printf("Unable to save InterfaceType: %v", err)
		http.Error(out, "Unable to save InterfaceType", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}
func DeleteInterfaceType(out http.ResponseWriter, in *http.Request) {
	var input InterfaceTypes.Model
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete InterfaceType: %v", err)
		http.Error(out, "Unable to delete InterfaceType", http.StatusInternalServerError)
		return

	}
	input.ID = id

	err = input.Delete()
	if err != nil {
		log.Printf("Unable to delete InterfaceType: %v", err)
		http.Error(out, "Unable to delete InterfaceType", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveDeviceType(out http.ResponseWriter, in *http.Request) {
	var input DeviceTypes.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save DeviceType: %v", err)
		http.Error(out, "Unable to save DeviceType", http.StatusInternalServerError)
		return
	}
	err = input.Save()
	if err != nil {
		log.Printf("Unable to save DeviceType: %v", err)
		http.Error(out, "Unable to save DeviceType", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}
func DeleteDeviceType(out http.ResponseWriter, in *http.Request) {
	var input DeviceTypes.Model
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete DeviceType: %v", err)
		http.Error(out, "Unable to delete DeviceType", http.StatusInternalServerError)
		return

	}
	input.ID = id

	err = input.Delete()
	if err != nil {
		log.Printf("Unable to delete DeviceType: %v", err)
		http.Error(out, "Unable to delete DeviceType", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveVLAN(out http.ResponseWriter, in *http.Request) {
	var input VLANs.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save VLAN: %v", err)
		http.Error(out, "Unable to save VLAN", http.StatusInternalServerError)
		return
	}
	err = input.Save()
	if err != nil {
		log.Printf("Unable to save VLAN: %v", err)
		http.Error(out, "Unable to save VLAN", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}
func DeleteVLAN(out http.ResponseWriter, in *http.Request) {
	var input VLANs.Model
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete VLAN: %v", err)
		http.Error(out, "Unable to delete VLAN", http.StatusInternalServerError)
		return

	}
	input.ID = id

	err = input.Delete()
	if err != nil {
		log.Printf("Unable to delete VLAN: %v", err)
		http.Error(out, "Unable to delete VLAN", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

func SaveUser(out http.ResponseWriter, in *http.Request) {
	var input Users.Model
	err := formats.ReadJSON(in, &input)
	if err != nil {
		log.Printf("Unable to save User: %v", err)
		http.Error(out, "Unable to save User", http.StatusInternalServerError)
		return
	}
	err = input.Save()
	if err != nil {
		log.Printf("Unable to save User: %v", err)
		http.Error(out, "Unable to save User", http.StatusInternalServerError)
		return

	}
	formats.PublishJSON(input, out, in)
}
func DeleteUser(out http.ResponseWriter, in *http.Request) {
	var input Users.Model
	id, err := strconv.ParseInt(mux.Vars(in)["ID"], 10, 64)
	if err != nil {
		log.Printf("Unable to delete User: %v", err)
		http.Error(out, "Unable to delete User", http.StatusInternalServerError)
		return

	}
	input.ID = id

	err = input.Delete()
	if err != nil {
		log.Printf("Unable to delete User: %v", err)
		http.Error(out, "Unable to delete User", http.StatusInternalServerError)
		return
	}
	formats.PublishJSON(input, out, in)
}

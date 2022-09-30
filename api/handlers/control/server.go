package control

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"facemasq/lib/db"
	"facemasq/lib/files"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/lib/netscan"
	"facemasq/lib/network"
	"facemasq/lib/portscan"
	"facemasq/lib/utils"
	"facemasq/models"

	"github.com/gorilla/mux"
)

func Exit(out http.ResponseWriter, in *http.Request) {
	logging.Processln("Remote exit invoked")
	os.Exit(0)
}

func Static(out http.ResponseWriter, in *http.Request) {
	file := mux.Vars(in)["filename"]
	if files.FileExists("../web/" + file) {
		out.Header().Set("Content-Type", mime.TypeByExtension(strings.TrimRight(filepath.Ext(file), "/")))
		http.ServeFile(out, in, "../web/"+file)
	} else {
		out.Header().Set("Content-Type", mime.TypeByExtension(strings.TrimRight(filepath.Ext(file), "/")))
		http.ServeFile(out, in, "../web/index.html")
	}
}

func Status(out http.ResponseWriter, in *http.Request) {
	var settings models.Meta
	details := make(map[string]string)

	details["NetScanFrequency"] = netscan.Frequency.String()
	details["PortScanActive"] = utils.Ternary(portscan.PortScan, "true", "false").(string)
	details["DBEngine"] = db.DBEngine
	details["NetMask"] = network.Target
	details["FormatHostnames"] = ""

	err := db.Conn.NewSelect().Model(&settings).Where(`name = 'formatHostnames' AND user_id IS NULL`).Scan(db.Context)
	if err != nil {
		logging.Errorf("error getting settings: %v", err)
		http.Error(out, "Unable to retrieve data", http.StatusInternalServerError)
		return
	}
	details["FormatHostnames"] = settings.Value

	formats.WriteJSONResponse(details, out, in)
}

package control

import (
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"facemasq/lib/db"
	"facemasq/lib/files"
	"facemasq/lib/formats"
	"facemasq/lib/netscan"
	"facemasq/lib/network"
	"facemasq/lib/portscan"
	"facemasq/lib/utils"

	"github.com/gorilla/mux"
)

func Exit(out http.ResponseWriter, in *http.Request) {
	log.Println("Remote exit invoked")
	os.Exit(0)
}

func Static(out http.ResponseWriter, in *http.Request) {
	file := mux.Vars(in)["filename"]
	if files.FileExists("../ui/" + file) {
		out.Header().Set("Content-Type", mime.TypeByExtension(strings.TrimRight(filepath.Ext(file), "/")))
		http.ServeFile(out, in, "../ui/"+file)
	} else {
		out.Header().Set("Content-Type", mime.TypeByExtension(strings.TrimRight(filepath.Ext(file), "/")))
		http.ServeFile(out, in, "../ui/index.html")
	}
}

func Status(out http.ResponseWriter, in *http.Request) {
	details := make(map[string]string)

	details["NetScanFrequency"] = netscan.Frequency.String()
	details["PortScanActive"] = utils.Ternary(portscan.PortScan, "true", "false").(string)
	details["DBEngine"] = db.DBEngine
	details["NetMask"] = network.Target

	formats.PublishJSON(details, out, in)
}

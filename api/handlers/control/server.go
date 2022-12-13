package control

import (
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"facemasq/lib/db"
	"facemasq/lib/files"
	"facemasq/lib/formats"
	"facemasq/lib/logging"
	"facemasq/lib/network"
	"facemasq/lib/scans/iprange"
	scanPorts "facemasq/lib/scans/port"
	"facemasq/lib/utils"
	"facemasq/models"

	"github.com/uptrace/bunrouter"
)

func Exit(out http.ResponseWriter, in bunrouter.Request) (err error) {
	logging.Info("Remote exit invoked")
	os.Exit(0)
	return
}

func Static(out http.ResponseWriter, in bunrouter.Request) (err error) {
	file := in.Params().ByName("filename")
	if strings.Contains(file, "i18n") {
		i18nDir, _ := files.GetDir("i18n")
		file = path.Join(i18nDir, strings.Replace(file, "i18n/", "", -1))
		logging.Info(file)
		out.Header().Set("Content-Type", mime.TypeByExtension(strings.TrimRight(filepath.Ext(file), "/")))
		http.ServeFile(out, in.Request, file)
	} else {
		var webDir string
		webDir, err = files.GetDir("dist/web")
		if err != nil {
			webDir, err = files.GetDir("web")
			if err != nil {
				return
			}
		}
		if files.FileExists(path.Join(webDir, file)) {
			out.Header().Set("Content-Type", mime.TypeByExtension(strings.TrimRight(filepath.Ext(file), "/")))
			http.ServeFile(out, in.Request, path.Join(webDir, file))
		} else {
			out.Header().Set("Content-Type", mime.TypeByExtension(strings.TrimRight(filepath.Ext(file), "/")))
			http.ServeFile(out, in.Request, path.Join(webDir, "index.html"))
		}

	}
	return
}

func State(out http.ResponseWriter, in bunrouter.Request) (err error) {
	var settings models.Meta
	details := make(map[string]string)

	details["NetScanFrequency"] = iprange.Frequency.String()
	details["PortScanActive"] = utils.Ternary(scanPorts.PortScan, "true", "false").(string)
	details["DBEngine"] = db.DBEngine
	details["NetMask"] = network.Target
	details["FormatHostnames"] = ""

	err = db.Conn.NewSelect().Model(&settings).Where(`name = 'formatHostnames' AND user_id IS NULL`).Scan(db.Context)
	if err != nil {
		return
	}
	details["FormatHostnames"] = settings.Value

	formats.WriteJSONResponse(details, out, in)
	return
}

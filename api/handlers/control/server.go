package control

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jsnfwlr/facemasq/api/lib/files"
)

func Exit(out http.ResponseWriter, in *http.Request) {
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

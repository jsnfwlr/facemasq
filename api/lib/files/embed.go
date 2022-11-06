package files

import (
	"embed"
	"io/fs"

	"facemasq/lib/logging"
)

var content embed.FS

func GetEmbeddedFileContents(name string) ([]byte, error) {
	logging.Debug2("Get %s", name)
	return content.ReadFile(name)
}

func GetEmbeddedFile(name string) (fs.File, error) {
	logging.Debug2("Get %s", name)
	return content.Open(name)
}

func GetEmbeddedFileSystem() embed.FS {
	return content
}

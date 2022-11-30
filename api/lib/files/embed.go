package files

import (
	"embed"
	"io/fs"
)

var content embed.FS

func GetEmbeddedFileContents(name string) ([]byte, error) {
	return content.ReadFile(name)
}

func GetEmbeddedFile(name string) (fs.File, error) {
	return content.Open(name)
}

func GetEmbeddedFileSystem() embed.FS {
	return content
}

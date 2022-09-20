package files

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed templates/*

var content embed.FS

func GetEmbeddedFileContents(name string) ([]byte, error) {
	fmt.Printf("Get %s", name)
	return content.ReadFile(name)
}

func GetEmbeddedFile(name string) (fs.File, error) {
	fmt.Printf("Get %s", name)
	return content.Open(name)
}

func GetEmbeddedFileSystem() embed.FS {
	return content
}

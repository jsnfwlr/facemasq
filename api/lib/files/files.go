package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"facemasq/lib/utils"
)

var mode = "detect"

func WriteOut(fileLocation, content string) (err error) {
	var fileHandler *os.File
	if _, err = os.Stat(fileLocation); err == nil {
		err = os.Remove(fileLocation)
		if err != nil {
			return
		}
		fileHandler, _ = os.Create(fileLocation)

	} else if os.IsNotExist(err) {
		fileHandler, err = os.Create(fileLocation)
		if err != nil {
			return
		}
	}
	_, err = fileHandler.WriteString(content)
	return
}

func FileExists(file string) (exists bool) {
	exists = false
	_, err := os.Stat(file)
	if err == nil {
		exists = true
	} else if os.IsNotExist(err) {
		exists = false
	}
	return
}

func Copy(src, dst string) (size int64, err error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return
	}

	if !sourceFileStat.Mode().IsRegular() {
		err = fmt.Errorf("%s is not a regular file", src)
		return
	}

	source, err := os.Open(src)
	if err != nil {
		return
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return
	}
	defer destination.Close()
	size, err = io.Copy(destination, source)
	return
}

func GetAppRoot() (rootDir string, err error) {
	if mode == "forced" {
		rootDir = "/app"
		return
	}
	rootDir, err = os.Getwd()
	if err != nil {
		return
	}
	return
}

func GetDir(which string) (dir string, err error) {
	var rootDir string
	rootDir, err = GetAppRoot()
	if err != nil {
		return
	}
	lowerWhich := strings.ToLower(which)
	switch lowerWhich {
	default:
		position := utils.Ternary(strings.Contains(rootDir, "api"), strings.Index(rootDir, "api"), 0).(int)
		dir = utils.Ternary((rootDir == "/app"), fmt.Sprintf("%[1]c%[2]s", os.PathSeparator, lowerWhich), fmt.Sprintf("%[1]s%[2]s", rootDir[0:position], lowerWhich)).(string)
	}
	if !FileExists(dir) {
		err = fmt.Errorf("could not find `%s`", dir)
	}
	return
}

func RemoveGlob(path string) (err error) {
	contents, err := filepath.Glob(path)
	if err != nil {
		return
	}
	for _, item := range contents {
		err = os.RemoveAll(item)
		if err != nil {
			return
		}
	}
	return
}

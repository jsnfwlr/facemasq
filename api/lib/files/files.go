package files

import (
	"fmt"
	"io"
	"os"
)

func WriteOut(fileLocation, content string) (err error) {
	var fileHandler *os.File
	if _, err = os.Stat(fileLocation); err == nil {
		os.Remove(fileLocation)
		fileHandler, err = os.Create(fileLocation)
		if err != nil {
			return
		}

	} else if os.IsNotExist(err) {
		fileHandler, err = os.Create(fileLocation)
		if err != nil {
			return
		}
	}
	fileHandler.WriteString(content)
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

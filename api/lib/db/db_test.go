package db

import (
	"os"
	"testing"
)

func TestDBInit(test *testing.T) {
	os.Remove("../../../data/test.sqlite")
	err := Connect("../../../data/", "test.sqlite")
	if err != nil {
		test.Error(err)
	}
}

package utils

import (
	"flag"
	"os"
	"strings"
)

func IsTest() (isTest bool) {
	isTest = strings.HasSuffix(os.Args[0], ".test")
	isTest = isTest || flag.Lookup("test.v") != nil
	return
}

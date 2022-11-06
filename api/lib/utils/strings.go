package utils

import (
	"regexp"
	"strings"
)

func StringContainsAny(haystack string, needles []string) bool {
	for i := range needles {
		if strings.Contains(haystack, needles[i]) {
			return true
		}
	}
	return false
}

func IsFormatString(haystack string) bool {
	needles := []*regexp.Regexp{
		regexp.MustCompile(`%[+#\.\d\*%]{0,3}[vTtbcdoOqxXUeEfFgGsp]{0,1}`),
		regexp.MustCompile(`%[+#\.\d\*]{0,3}\[\d+\][vTtbcdoOqxXUeEfFgGsp]`),
	}
	for n := range needles {
		if needles[n].MatchString(haystack) {
			return true
		}
	}
	return false
}

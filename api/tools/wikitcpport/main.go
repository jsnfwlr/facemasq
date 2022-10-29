package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var err error
	var response *http.Response
	var lines []string
	url := "https://en.wikipedia.org/w/index.php?title=List_of_TCP_and_UDP_port_numbers&action=edit"
	response, err = http.Get(url)
	if err != nil {
		return
	}
	buf := new(strings.Builder)
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		return
	}

	content := buf.String()
	lines, err = sortLines(content)
	if err != nil {
		panic(err)
	}

	err = parseLines(lines)
	if err != nil {
		panic(err)
	}
}

type searchReplace struct {
	name    string
	find    *regexp.Regexp
	replace string
}

func getRegex(mode string) (newline, multirow *regexp.Regexp, reggies []searchReplace) {
	switch mode {
	case "sort":
		reggies = []searchReplace{
			{name: "single-multi", find: regexp.MustCompile(`^\|\s{0,10}\d+\s{0,10}\|\| colspan=\d+`), replace: ""},
			{name: "single-single", find: regexp.MustCompile(`^\|\s{0,10}\d+\s{0,10}\|\| \{\{`), replace: ""},
			{name: "single-blank", find: regexp.MustCompile(`^\|\s{0,10}\d+\s{0,10}\|\|\s{1,10}`), replace: ""},
			{name: "single-what", find: regexp.MustCompile(`^\|\s{0,10}\d+\s{0,10}\|`), replace: ""},
			{name: "multi-multi", find: regexp.MustCompile(`^\|\s{0,10}\d+[^\d]\d+\s{0,10}\|\| colspan=\d+`), replace: ""},
			{name: "multi-single", find: regexp.MustCompile(`^\|\s{0,10}\d+[^\d]\d+\s{0,10}\|\|\s{0,10}\{\{`), replace: ""},
			{name: "multi-blank", find: regexp.MustCompile(`^\|\s{0,10}\d+[^\d]\d+\s{0,10}\|\|\s{1,10}`), replace: ""},
			{name: "multi-row", find: regexp.MustCompile(`^\|\s{0,10}rowspan="?(\d+)"?\s{0,10}\|`), replace: ""},
		}
		newline = regexp.MustCompile(`\|-`)
		multirow = reggies[7].find
	case "sanitise":
		reggies = []searchReplace{
			{name: "col-span", find: regexp.MustCompile(`\|\|\s{0,10}colspan=2([^\|]+)\|\|`), replace: "|| $1 || $1 ||"},
			{name: "split-col", find: regexp.MustCompile(`\|\|\|`), replace: "|| ? ||"},
			{name: "simple-decode", find: regexp.MustCompile(`&lt;code>([^>]+)&lt;/code>`), replace: "$1"},
			{name: "simple-deabbr", find: regexp.MustCompile(`&lt;abbr>([^>]+)&lt;/abbr>`), replace: "$1"},
			{name: "simple-href", find: regexp.MustCompile(`\[https?://[^\s\]]+\s([^\]]+)\]`), replace: "$1"},
			{name: "complex-delink", find: regexp.MustCompile(`\[\[[^\]]+\|([^\]]+)\]\]`), replace: "$1"},
			{name: "simple-delink", find: regexp.MustCompile(`\[\[([^\]]+)\]\]`), replace: "$1"},
			{name: "simple-destyle", find: regexp.MustCompile(`''([^']+)''`), replace: "$1"},

			{name: "simple-tag", find: regexp.MustCompile(`\{\{([^\s\}]+)\}\}`), replace: "$1"},
			{name: "complex-deaudit", find: regexp.MustCompile(`\{\{[^\]]+\|date[^\]]+\}\}`), replace: ""},
			{name: "complex-decitation", find: regexp.MustCompile(`\{\{cite[^\}]+\}\}`), replace: ""},
			{name: "simple-deasof", find: regexp.MustCompile(`\{\{As of\|\d{4}\}\}`), replace: ""},

			{name: "simple-named-dereference", find: regexp.MustCompile(`&lt;ref name\s?=\s?"[^>]+"\s?/>`), replace: ""},
			{name: "simple-unquoted-dereference", find: regexp.MustCompile(`&lt;ref name\s?=\s?[^>]+\s?/>`), replace: ""},
			{name: "simple-group-dereference", find: regexp.MustCompile(`&lt;ref group\s?=\s?"[^>]+"\s?/>`), replace: ""},
			{name: "simple-groupunquoted-dereference", find: regexp.MustCompile(`&lt;ref group\s?=\s?[^>]+\s?/>`), replace: ""},
			{name: "complex-named-dereference", find: regexp.MustCompile(`&lt;ref name\s?=\s?"[^>]+">[^>]+&lt;/ref>`), replace: ""},
			{name: "complex-group-dereference", find: regexp.MustCompile(`&lt;ref group\s?=\s?"[^>]+">[^>]+&lt;/ref>`), replace: ""},
			{name: "complex-unquoted-dereference", find: regexp.MustCompile(`&lt;ref name\s?=\s?[^>]+>[^>]+&lt;/ref>`), replace: ""},
			{name: "complex-unamed-dereference", find: regexp.MustCompile(`&lt;ref>[^>]+&lt;/ref>`), replace: ""},

			{name: "empty-named-dereference", find: regexp.MustCompile(`&lt;ref name\s?=\s?"[^>]+">+&lt;/ref>`), replace: ""},
			{name: "empty-unquoted-dereference", find: regexp.MustCompile(`&lt;ref name\s?=\s?[^>]+>+&lt;/ref>`), replace: ""},
			{name: "empty-unamed-dereference", find: regexp.MustCompile(`&lt;ref>&lt;/ref>`), replace: ""},
			{name: "broken-named-dereference", find: regexp.MustCompile(`&lt;ref name\s?=\s?"[^>]+"\s?>$`), replace: ""},
			{name: "broken-unnamed-dereference", find: regexp.MustCompile(`&lt;ref\s?>$`), replace: ""},
			{name: "broken-unnamed-dereference", find: regexp.MustCompile(`&lt;ref\s?>[^>]+\|`), replace: "|"},
			{name: "broken-unnamed-dereference", find: regexp.MustCompile(`&lt;ref\s?>\|`), replace: "|"},
			{name: "broken-unnamed-dereference", find: regexp.MustCompile(`&lt;ref>\? \-\->`), replace: ""},
			{name: "cleanup", find: regexp.MustCompile(`^\|\s{0,10}(\d+)\s{0,10}\|\|`), replace: "| $1 "},
			{name: "empty-col", find: regexp.MustCompile(`\|\s{1,10}\|`), replace: "| ? |"},
			{name: "broken-col", find: regexp.MustCompile(`^\|\s{1,10}(\d+)\s{0,10}([^\|])`), replace: "| $1 || $2"},
			{name: "broken-col", find: regexp.MustCompile(`Maybe\|Assigned`), replace: "Assigned"},
			{name: "broken-col", find: regexp.MustCompile(`N/A\|Reserved`), replace: "Reserved"},
			{name: "broken-col", find: regexp.MustCompile(`n/a\|Reserved`), replace: "Reserved"},
			{name: "rowspan", find: regexp.MustCompile(`-->\|`), replace: " ||"},
			{name: "rowspan", find: regexp.MustCompile(`--> \|`), replace: " ||"},
			{name: "simple-decomment", find: regexp.MustCompile(`&lt;!--[^>\d]+-->`), replace: ""},
		}
		//

	case "normalise":
		reggies = []searchReplace{
			{name: "rowspan", find: regexp.MustCompile(`^\|\s{0,10}rowspan="?(\d{1,10})"?\s{0,10}\|(.*)`), replace: ""},
			{name: "range", find: regexp.MustCompile(`^\|\s{0,10}(\d+)[^\d](\d+)\s{0,10}\|\|(.*)`), replace: ""},
			{name: "colspan", find: regexp.MustCompile(`colspan=(\d+)\s{0,10}(\{\{[^\}]+\}\}|\w+[^\|]+)`), replace: ""},
		}
	}
	return
}

func parseLines(lines []string) (err error) {
	var normalisedLines []string
	var first, last, cols int
	// var rows int

	var rest string
	_, _, santise := getRegex("sanitise")
	_, _, reggies := getRegex("normalise")

	for l := range lines {
		matched := false
		for r := range reggies {
			if reggies[r].find.MatchString(lines[l]) {
				matched = true

				switch reggies[r].name {
				case "rowspan":
					parts := reggies[r].find.FindStringSubmatch(lines[l])
					rest = parts[2]
					newRows := strings.Split(rest, "&lt;!--")
					for n := range newRows {
						if n > 0 {
							newLine := fmt.Sprintf("| %s", newRows[n])
							normalisedLines = append(normalisedLines, newLine)
						}
					}

				case "range":
					parts := reggies[r].find.FindStringSubmatch(lines[l])
					first, err = strconv.Atoi(parts[1])
					if err != nil {
						return
					}
					last, err = strconv.Atoi(parts[2])
					if err != nil {
						return
					}
					rest := parts[3]
					for i := first; i <= last; i++ {
						newLine := fmt.Sprintf("| %d ||%s", i, rest)
						normalisedLines = append(normalisedLines, newLine)
					}
				case "colspan":
					parts := reggies[r].find.FindStringSubmatch(lines[l])
					cols, err = strconv.Atoi(parts[1])
					rest = parts[2]
					repeats := ""
					for i := 0; i < cols; i++ {
						repeats = fmt.Sprintf("%s%s || ", repeats, rest)
					}
					newLine := strings.Replace(lines[l], parts[0], repeats, 1)
					normalisedLines = append(normalisedLines, newLine)
				}
			}
		}
		if !matched {
			normalisedLines = append(normalisedLines, lines[l])
		}
	}

	for n := range normalisedLines {
		for s := range santise {
			normalisedLines[n] = santise[s].find.ReplaceAllString(normalisedLines[n], santise[s].replace)
		}
		fmt.Println(normalisedLines[n])
	}
	return
}

func sortLines(content string) (results []string, err error) {
	var record []string
	var wasNew bool

	output := false
	nextRows := 0

	newline, multirow, reggies := getRegex("sort")

	lines := strings.Split(content, "\n")
	for i := 0; i < len(lines); i++ {
		if lines[i] == "|}" {

			if output {

				line := strings.Join(record, "")
				for r := range reggies {
					if reggies[r].find.MatchString(line) {
						break
					}
				}
				if line != "" {
					results = append(results, line)
				}
			}
			output = false
			record = nil
		}

		if output {

			if newline.MatchString(lines[i]) {
				if nextRows == 0 && !wasNew {
					line := strings.Join(record, "")
					for r := range reggies {
						if reggies[r].find.MatchString(line) {
							break
						}
					}

					if line != "" {
						results = append(results, line)
					}
					record = nil
					continue
				} else if !wasNew {
					nextRows--
				}
				wasNew = true
			} else {
				wasNew = false
			}

			if multirow.MatchString(lines[i]) {
				matches := multirow.FindStringSubmatch(lines[i])
				nextRows, err = strconv.Atoi(matches[1])
				if err != nil {
					err = fmt.Errorf("%v - %v", matches, err)
					return
				}
				nextRows -= 1 // the current row is the first in the list
			}
			if string(lines[i][0]) != "!" {
				record = append(record, lines[i])
			}
		}

		if lines[i] == "|+ Well-known ports" || lines[i] == "|+ Dynamic, private or ephemeral ports" || lines[i] == "|+ Registered ports" {
			output = true
		}
	}
	return
}

package logging

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

var Want string
var timeFormat = "2006-01-02T15:04:05"

func TestLoggingLevels(t *testing.T) {
	AllowTestLogging = true
	SetLevel(zerolog.TraceLevel)

	firstline := 56
	tests := []struct {
		level string
		line  int

		fmt      string
		arg1     string
		arg2     string
		expected string
	}{
		{level: "trace", line: firstline + 0, fmt: "%s => logged as %s", arg1: "Trace", arg2: "Trace", expected: `{"level":"%[1]s","time":"%[2]s","caller":"/home/phalacee/Projects/apps/facemasq/api/lib/logging/logging_test.go:%[3]d","message":"%[4]s"}`},
		{level: "debug", line: firstline + 2, fmt: "%s => logged as %s", arg1: "Debug", arg2: "Debug", expected: `{"level":"%[1]s","time":"%[2]s","caller":"/home/phalacee/Projects/apps/facemasq/api/lib/logging/logging_test.go:%[3]d","message":"%[4]s"}`},
		{level: "info", line: firstline + 4, fmt: "%s => logged as %s", arg1: "Info", arg2: "Info", expected: `{"level":"%[1]s","time":"%[2]s","caller":"/home/phalacee/Projects/apps/facemasq/api/lib/logging/logging_test.go:%[3]d","message":"%[4]s"}`},
		{level: "warn", line: firstline + 6, fmt: "%s => logged as %s", arg1: "Warning", arg2: "Warning", expected: `{"level":"%[1]s","time":"%[2]s","caller":"/home/phalacee/Projects/apps/facemasq/api/lib/logging/logging_test.go:%[3]d","message":"%[4]s"}`},
		{level: "error", line: firstline + 8, fmt: "%s => logged as %s", arg1: "Error", arg2: "Error", expected: `{"level":"%[1]s","error":"%[4]s","time":"%[2]s","caller":"/home/phalacee/Projects/apps/facemasq/api/lib/logging/logging_test.go:%[3]d"}`},
		{level: "panic", line: firstline + 10, fmt: "%s => logged as %s", arg1: "Panic", arg2: "Panic", expected: `{"level":"%[1]s","error":"%[4]s","time":"%[2]s","caller":"/home/phalacee/Projects/apps/facemasq/api/lib/logging/logging_test.go:%[3]d"}`},
		// {level: "fatal", line: firstline + 12, fmt: "%s => logged as %s", arg1: "Fatal", arg2: "Fatal", expected: `{"level":"%[1]s","error":"%[4]s","time":"%[2]s","caller":"/home/phalacee/Projects/apps/facemasq/api/lib/logging/logging_test.go:%[3]d"}`},
		// {Level: "trace", line: 30, fmt: "%s => logged as %s", arg1: "Trace", arg2: "Trace"},
	}

	have := bytes.Buffer{}
	Init(&have, &have, timeFormat)

	defer func(have *bytes.Buffer) {
		if r := recover(); r != nil {
			if strings.TrimSuffix(have.String(), "\n") != Want {
				t.Errorf("Error with %s:\nhave: '%s'\nwant: '%s'\n", "panic", strings.TrimSuffix(have.String(), "\n"), Want)
			}
		}
	}(&have)

	for i := range tests {
		have = bytes.Buffer{}
		Want = fmt.Sprintf(tests[i].expected, tests[i].level, time.Now().Format(timeFormat), tests[i].line, fmt.Sprintf(tests[i].fmt, tests[i].arg1, tests[i].arg2))
		switch tests[i].arg1 {
		case "Trace":
			Trace(tests[i].fmt, tests[i].arg1, tests[i].arg2)
		case "Debug":
			Debug(tests[i].fmt, tests[i].arg1, tests[i].arg2)
		case "Info":
			Info(tests[i].fmt, tests[i].arg1, tests[i].arg2)
		case "Warning":
			Warning(tests[i].fmt, tests[i].arg1, tests[i].arg2)
		case "Error":
			Error(tests[i].fmt, tests[i].arg1, tests[i].arg2)
		case "Panic":
			Panic(fmt.Sprintf(tests[i].fmt, tests[i].arg1, tests[i].arg2))
		case "Fatal":
			Fatal(tests[i].fmt, tests[i].arg1, tests[i].arg2)
		}

		if strings.TrimSuffix(have.String(), "\n") != Want {
			t.Errorf("Error with %s:\nhave: '%s'\nwant: '%s'\n", tests[i].level, strings.TrimSuffix(have.String(), "\n"), Want)
		}
	}
}

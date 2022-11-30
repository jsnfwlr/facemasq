package logging

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"

	"facemasq/lib/utils"
)

var Verbosity int

var logger = New("", "")

const (
	ALWAYS                 = iota // 1
	WARN                          // 2
	NOTICE                        // 3
	INFORM                        // 4
	DEBUG1                        // 5
	DEBUG2                        // 6
	DEBUG3                        // 7
	SYSTEM                 = ALWAYS
	ERROR                  = ALWAYS
	FATAL                  = ALWAYS
	PANIC                  = ALWAYS
	defaultTemplate        = "[{{ .Class }}] {{ .Timestamp }} | {{ .TLA }} | {{ .Location }} | {{ .Message }}\n"
	defaultTimestampFormat = time.RFC3339
)

type Logger interface {
	SetTemplate(string) error
	SetTimestampFormat(string) error
	SetStdout(io.WriteCloser)
	SetStderr(io.WriteCloser)
	GetTemplate() *template.Template
	GetTimestampFormat() string
	GetStdout() io.WriteCloser
	GetStderr() io.WriteCloser
	Debug1(interface{}, ...interface{})
	Debug2(interface{}, ...interface{})
	Debug3(interface{}, ...interface{})
	Warning(interface{}, ...interface{})
	Notice(interface{}, ...interface{})
	Info(interface{}, ...interface{})

	System(interface{}, ...interface{})
	Error(interface{}, ...interface{})
	Fatal(interface{}, ...interface{})
	Panic(interface{}, ...interface{})
}

type Default struct {
	template        *template.Template
	timestampFormat string
	stdOut          io.WriteCloser
	stdErr          io.WriteCloser
}

type messenger struct {
	Class     string
	Timestamp string
	TLA       string
	Location  string
	Message   string
}

func init() {
	var err error
	Verbosity, err = strconv.Atoi(os.Getenv("VERBOSE"))
	if err != nil {
		Verbosity = 0
	}
	fmt.Printf("Log Verbosity set to %d\n", Verbosity)
}

func New(tmpl, timestampFormat string) (logger Logger) {
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}
	logger = &Default{
		template:        setTemplate(tmpl),
		timestampFormat: timestampFormat,
		stdOut:          os.Stdout,
		stdErr:          os.Stderr,
	}
	return
}

func setTemplate(tmpl string) (parsed *template.Template) {
	var err error
	if tmpl == "" {
		tmpl = defaultTemplate
	}
	parsed, err = template.New("set").Parse(defaultTemplate)
	if err != nil {
		panic(err)
	}
	return
}

func prepareMessage(arg0 interface{}, args ...interface{}) (msg string) {
	if reflect.TypeOf(arg0).String() == "string" && utils.IsFormatString(arg0.(string)) {
		msg = fmt.Sprintf(arg0.(string), args...)
		return
	}
	newArgs := args
	newArgs = nil
	newArgs = append(newArgs, arg0)
	newArgs = append(newArgs, args...)
	msg = fmt.Sprint(newArgs...)
	return
}

func (logger *Default) SetTemplate(tmpl string) (err error) {
	if tmpl == "" {
		tmpl = defaultTemplate
	}
	logger.template, err = template.New("set").Parse(tmpl)
	return
}

func (logger *Default) SetTimestampFormat(timestampFormat string) (err error) {
	logger.timestampFormat = timestampFormat
	if logger.timestampFormat == "" {
		logger.timestampFormat = defaultTimestampFormat
	}
	return
}

func (logger *Default) SetStdout(stdOut io.WriteCloser) {
	logger.stdOut = stdOut
}

func (logger *Default) SetStderr(stdErr io.WriteCloser) {
	logger.stdErr = stdErr
}

func (logger *Default) GetTemplate() *template.Template {
	return logger.template
}

func (logger *Default) GetTimestampFormat() string {
	return logger.timestampFormat
}

func (logger *Default) GetStdout() io.WriteCloser {
	return logger.stdOut
}

func (logger *Default) GetStderr() io.WriteCloser {
	return logger.stdErr
}

func (logger *Default) Debug(arg0 interface{}, args ...interface{}) {
	if Verbosity >= DEBUG1 {
		logger.output("debug  ", prepareMessage(arg0, args...), 2)
	}
}

func (logger *Default) Debug1(arg0 interface{}, args ...interface{}) {
	if Verbosity >= DEBUG1 {
		logger.output("debug  ", prepareMessage(arg0, args...), 2)
	}
}

func (logger *Default) Debug2(arg0 interface{}, args ...interface{}) {
	if Verbosity >= DEBUG2 {
		logger.output("debug  ", prepareMessage(arg0, args...), 2)
	}
}

func (logger *Default) Debug3(arg0 interface{}, args ...interface{}) {
	if Verbosity >= DEBUG3 {
		logger.output("debug  ", prepareMessage(arg0, args...), 2)
	}
}

func (logger *Default) Info(arg0 interface{}, args ...interface{}) {
	if Verbosity >= INFORM {
		logger.output("info   ", prepareMessage(arg0, args...), 2)
	}
}

func (logger *Default) Warning(arg0 interface{}, args ...interface{}) {
	if Verbosity >= WARN {
		logger.output("warning", prepareMessage(arg0, args...), 2)
	}
}

func (logger *Default) Notice(arg0 interface{}, args ...interface{}) {
	if Verbosity >= INFORM {
		logger.output("notice ", prepareMessage(arg0, args...), 2)
	}
}

func (logger *Default) System(arg0 interface{}, args ...interface{}) {
	if Verbosity >= SYSTEM {
		logger.output("system ", prepareMessage(arg0, args...), 2)
	}
}

func (logger *Default) Error(arg0 interface{}, args ...interface{}) {
	if Verbosity >= ERROR {
		logger.output("error  ", prepareMessage(arg0, args...), 2)
	}
}

func (logger *Default) Panic(arg0 interface{}, args ...interface{}) {
	if Verbosity >= PANIC {
		logger.output("panic  ", prepareMessage(arg0, args...), 2)
	}
	panic("...")
}

func (logger *Default) Fatal(arg0 interface{}, args ...interface{}) {
	if Verbosity >= FATAL {
		logger.output("fatal  ", prepareMessage(arg0, args...), 2)
	}
	os.Exit(1)
}

func (logger *Default) output(class, message string, depth int) {
	if logger.timestampFormat == "" {
		logger.timestampFormat = defaultTimestampFormat
	}
	if logger.template.Name() != "set" {
		logger.template = template.Must(template.New("set").Parse(defaultTemplate))
	}
	output(class, message, depth, logger.timestampFormat, logger.template, logger.stdOut, logger.stdErr)

}

func output(class, message string, depth int, timestampFormat string, tmpl *template.Template, stdOut, errOut io.WriteCloser) {
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}
	if tmpl == nil {
		tmpl = template.Must(template.New("set").Parse(defaultTemplate))
	}

	_, file, line, _ := runtime.Caller(depth)

	root, err := getAppRoot()
	if err != nil {
		root = ""
	}

	file = strings.Replace(file, root, "", -1)

	msg := messenger{
		Class:     class,
		Timestamp: time.Now().Format(timestampFormat),
		Location:  fmt.Sprintf("%s:%d", file, line),
		TLA:       "000",
		Message:   message,
	}
	switch class {
	case "error", "panic", "fatal":

		_ = tmpl.Execute(errOut, msg)
	default:
		_ = tmpl.Execute(stdOut, msg)
	}
}

func getAppRoot() (rootDir string, err error) {
	rootDir, err = os.Getwd()
	if err != nil {
		return
	}
	return
}

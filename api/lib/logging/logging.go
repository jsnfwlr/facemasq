package logging

import (
	"fmt"
	"io"
	"log/syslog"
	"os"
	"path"
	"reflect"

	"facemasq/lib/files"
	"facemasq/lib/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var stdOut, stdErr zerolog.Logger
var AllowTestLogging = true
var SysLog, FileLog bool

func init() {
	var sysOut *syslog.Writer
	var fileOut *os.File
	var err error
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000"

	conOut := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02T15:04:05.000"}
	conErr := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02T15:04:05.000"}
	SysLog = true
	if SysLog {
		sysOut, err = syslog.Dial("tcp", "192.168.0.44:514", syslog.LOG_EMERG|syslog.LOG_ERR|syslog.LOG_INFO|syslog.LOG_CRIT|syslog.LOG_WARNING|syslog.LOG_NOTICE|syslog.LOG_DEBUG, "faceMasq")
		if err != nil {
			panic(err)
		}
	}

	FileLog = true
	if FileLog {
		dir, err := files.GetDir("logs")
		if err != nil {
			panic(err)
		}
		fileOut, err = os.OpenFile(path.Join(dir, "facemasq.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			panic(err)
		}
	}

	stdOut = zerolog.New(zerolog.MultiLevelWriter(conOut, fileOut, sysOut)).With().Timestamp().CallerWithSkipFrameCount(utils.Ternary(utils.IsTest(), 3, 3).(int)).Logger()
	stdErr = zerolog.New(zerolog.MultiLevelWriter(conErr, fileOut, sysOut)).With().Timestamp().CallerWithSkipFrameCount(utils.Ternary(utils.IsTest(), 3, 3).(int)).Logger()

	SetLevel(zerolog.InfoLevel)

}

func SetLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)

}

func Init(writerOut io.Writer, writerErr io.Writer, timeFormat string) {
	zerolog.TimeFieldFormat = timeFormat
	stdOut = zerolog.New(writerOut).With().Timestamp().CallerWithSkipFrameCount(utils.Ternary(utils.IsTest(), 3, 0).(int)).Logger()
	stdErr = zerolog.New(writerErr).With().Timestamp().CallerWithSkipFrameCount(utils.Ternary(utils.IsTest(), 3, 0).(int)).Logger()
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

func Trace(arg0 interface{}, args ...interface{}) {
	if AllowTestLogging || !utils.IsTest() {
		stdOut.Trace().Msg(prepareMessage(arg0, args...))
	}
}

func Debug(arg0 interface{}, args ...interface{}) {
	stdOut.Debug().Msg(prepareMessage(arg0, args...))
}

func Info(arg0 interface{}, args ...interface{}) {
	if AllowTestLogging || !utils.IsTest() {
		stdOut.Info().Msg(prepareMessage(arg0, args...))
	}
}

func Warning(arg0 interface{}, args ...interface{}) {
	stdOut.Warn().Msg(prepareMessage(arg0, args...))
}

func Error(arg0 interface{}, args ...interface{}) {
	stdErr.Error().Err(errors.New(prepareMessage(arg0, args...))).Msg("")
}

func Fatal(arg0 interface{}, args ...interface{}) {
	stdErr.Fatal().Err(errors.New(prepareMessage(arg0, args...))).Msg("")
}

func Panic(arg0 interface{}, args ...interface{}) {
	stdErr.Panic().Err(errors.New(prepareMessage(arg0, args...))).Msg("")
}

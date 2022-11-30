package logging

import (
	"fmt"
	"log/syslog"
	"os"
	"reflect"

	"facemasq/lib/files"
	"facemasq/lib/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var stdOut, stdErr zerolog.Logger

func init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000"

	conOut := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02T15:04:05.000"}
	conErr := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02T15:04:05.000"}

	sysOut, err := syslog.Dial("tcp", "192.168.0.44:514", syslog.LOG_EMERG|syslog.LOG_ERR|syslog.LOG_INFO|syslog.LOG_CRIT|syslog.LOG_WARNING|syslog.LOG_NOTICE|syslog.LOG_DEBUG, "faceMasq")
	if err != nil {
		panic(err)
	}
	// plainOut := os.Stdout
	// plainErr := os.Stderr

	dir, err := files.GetDir("logs")
	if err != nil {
		panic(err)
	}

	outFile := fmt.Sprintf("%[2]s%[1]c%[3]s", os.PathSeparator, dir, "facemasq.log")
	fmt.Println(outFile)
	fileOut, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}

	multiOut := zerolog.MultiLevelWriter(conOut, fileOut, sysOut)
	multiErr := zerolog.MultiLevelWriter(conErr, fileOut, sysOut)

	stdOut = zerolog.New(multiOut).With().Timestamp().Caller().Logger()
	stdErr = zerolog.New(multiErr).With().Timestamp().Caller().Logger()
	SetLevel(zerolog.InfoLevel)

}

func SetLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)

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
	stdOut.Trace().Msg(prepareMessage(arg0, args...))
}

func Debug(arg0 interface{}, args ...interface{}) {
	stdOut.Debug().Msg(prepareMessage(arg0, args...))
}

func Info(arg0 interface{}, args ...interface{}) {
	stdOut.Info().Msg(prepareMessage(arg0, args...))
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

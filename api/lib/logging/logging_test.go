package logging

import (
	"testing"
)

func TestLogging(t *testing.T) {
	Trace("Logging a %s", "trace")
	Debug("Logging a %s", "debug")
	Info("Logging a %s", "info")
	Warning("Logging a %s", "warn")
	Error("Logging a %s", "err")
	Fatal("Logging a %s", "fatal")
	Panic("Logging a %s", "panic")
}

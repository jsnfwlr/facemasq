package logging

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var Verbosity int

func init() {
	var err error
	Verbosity, err = strconv.Atoi(os.Getenv("VERBOSE"))
	if err != nil {
		Verbosity = 0
	}
	Printf(0, "Log Verbosity set to %d", Verbosity)
}

func Printf(verbosity int, format string, args ...interface{}) {
	if Verbosity >= verbosity {
		fmt.Printf("[backend] %s | NFO | %s\n", time.Now().Format(time.RFC3339), fmt.Sprintf(format, args...))
	}
}

func Println(verbosity int, args ...interface{}) {
	if Verbosity >= verbosity {
		fmt.Printf("[backend] %s | NFO | %s\n", time.Now().Format(time.RFC3339), fmt.Sprint(args...))
	}
}

func Debugf(verbosity int, format string, args ...interface{}) {
	if Verbosity >= verbosity {
		fmt.Printf("[backend] %s | DBG | %s\n", time.Now().Format(time.RFC3339), fmt.Sprintf(format, args...))
	}
}

func Debugln(verbosity int, args ...interface{}) {
	if Verbosity >= verbosity {
		fmt.Printf("[backend] %s | DBG | %s\n", time.Now().Format(time.RFC3339), fmt.Sprint(args...))
	}
}

func Processf(format string, args ...interface{}) {
	fmt.Printf("[backend] %s | PRC | %s\n", time.Now().Format(time.RFC3339), fmt.Sprintf(format, args...))
}

func Processln(args ...interface{}) {
	fmt.Printf("[backend] %s | PRC | %s\n", time.Now().Format(time.RFC3339), fmt.Sprint(args...))
}

func Errorf(format string, args ...interface{}) {
	fmt.Printf("[backend] %s | ERR | %s\n", time.Now().Format(time.RFC3339), fmt.Sprintf(format, args...))
}

func Errorln(args ...interface{}) {
	fmt.Printf("[backend] %s | ERR | %s\n", time.Now().Format(time.RFC3339), fmt.Sprint(args...))
}

func Panicf(format string, args ...interface{}) {
	fmt.Printf("[backend] %s | PAN | %s", time.Now().Format(time.RFC3339), fmt.Sprintf(format, args...))
	panic("...")
}

func Panic(args ...interface{}) {
	fmt.Printf("[backend] %s | PAN | %s", time.Now().Format(time.RFC3339), fmt.Sprint(args...))
	panic("...")
}

func Fatalf(format string, args ...interface{}) {
	fmt.Printf("[backend] %s | FTL | %s", time.Now().Format(time.RFC3339), fmt.Sprintf(format, args...))
	os.Exit(1)
}

func Fatalln(args ...interface{}) {
	fmt.Printf("[backend] %s | FTL | %s", time.Now().Format(time.RFC3339), fmt.Sprint(args...))
	os.Exit(1)
}

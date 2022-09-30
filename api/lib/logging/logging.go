package logging

import "log"

var Verbosity int

func Printf(verbosity int, format string, args ...interface{}) {
	if Verbosity >= verbosity {
		log.Printf(format, args...)
	}
}

func Println(verbosity int, args ...interface{}) {
	if Verbosity >= verbosity {
		log.Println(args...)
	}
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Processf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Processln(args ...interface{}) {
	log.Println(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Errorln(args ...interface{}) {
	log.Println(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	log.Fatalln(args...)
}

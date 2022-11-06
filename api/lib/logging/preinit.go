package logging

import "os"

var preInitShown = false

func showPreInit() {
	if !preInitShown {
		preInitShown = true
		System("Calling logging function without initialisation")
	}
}

func Debug(arg0 interface{}, args ...interface{}) {
	if Verbosity >= DEBUG1 {
		if logger == nil {
			showPreInit()
			output("debug  ", prepareMessage(arg0, args...), 2, "", nil, os.Stdout, os.Stderr)
		} else {
			output("debug  ", prepareMessage(arg0, args...), 2, logger.GetTimestampFormat(), logger.GetTemplate(), logger.GetStdout(), logger.GetStderr())
		}
	}
}

func Debug1(arg0 interface{}, args ...interface{}) {
	if Verbosity >= DEBUG1 {
		if logger == nil {
			showPreInit()
			output("debug  ", prepareMessage(arg0, args...), 2, "", nil, os.Stdout, os.Stderr)
		} else {
			output("debug  ", prepareMessage(arg0, args...), 2, logger.GetTimestampFormat(), logger.GetTemplate(), logger.GetStdout(), logger.GetStderr())
		}
	}
}

func Debug2(arg0 interface{}, args ...interface{}) {
	if Verbosity >= DEBUG2 {
		if logger == nil {
			showPreInit()
			output("debug  ", prepareMessage(arg0, args...), 2, "", nil, os.Stdout, os.Stderr)
		} else {
			output("debug  ", prepareMessage(arg0, args...), 2, logger.GetTimestampFormat(), logger.GetTemplate(), logger.GetStdout(), logger.GetStderr())
		}
	}
}

func Debug3(arg0 interface{}, args ...interface{}) {
	if Verbosity >= DEBUG3 {
		if logger == nil {
			showPreInit()
			output("debug  ", prepareMessage(arg0, args...), 2, "", nil, os.Stdout, os.Stderr)
		} else {
			output("debug  ", prepareMessage(arg0, args...), 2, logger.GetTimestampFormat(), logger.GetTemplate(), logger.GetStdout(), logger.GetStderr())
		}
	}
}

func Info(arg0 interface{}, args ...interface{}) {

	if Verbosity >= INFORM {
		if logger == nil {
			showPreInit()
			output("info   ", prepareMessage(arg0, args...), 2, "", nil, os.Stdout, os.Stderr)
		} else {
			output("info   ", prepareMessage(arg0, args...), 2, logger.GetTimestampFormat(), logger.GetTemplate(), logger.GetStdout(), logger.GetStderr())
		}
	}
}

func Warning(arg0 interface{}, args ...interface{}) {

	if Verbosity >= WARN {
		if logger == nil {
			showPreInit()
			output("warning", prepareMessage(arg0, args...), 2, "", nil, os.Stdout, os.Stderr)
		} else {
			output("warning", prepareMessage(arg0, args...), 2, logger.GetTimestampFormat(), logger.GetTemplate(), logger.GetStdout(), logger.GetStderr())
		}
	}
}

func Notice(arg0 interface{}, args ...interface{}) {

	if Verbosity >= INFORM {
		if logger == nil {
			showPreInit()
			output("notice ", prepareMessage(arg0, args...), 2, "", nil, os.Stdout, os.Stderr)
		} else {
			output("notice ", prepareMessage(arg0, args...), 2, logger.GetTimestampFormat(), logger.GetTemplate(), logger.GetStdout(), logger.GetStderr())
		}
	}
}

func System(arg0 interface{}, args ...interface{}) {
	if Verbosity >= SYSTEM {
		if logger == nil {
			showPreInit()
			output("system ", prepareMessage(arg0, args...), 2, "", nil, os.Stdout, os.Stderr)
		} else {
			output("system ", prepareMessage(arg0, args...), 2, logger.GetTimestampFormat(), logger.GetTemplate(), logger.GetStdout(), logger.GetStderr())
		}
	}
}

func Error(arg0 interface{}, args ...interface{}) {

	if Verbosity >= ERROR {
		if logger == nil {
			showPreInit()
			output("error  ", prepareMessage(arg0, args...), 2, "", nil, os.Stdout, os.Stderr)
		} else {
			output("error  ", prepareMessage(arg0, args...), 2, logger.GetTimestampFormat(), logger.GetTemplate(), logger.GetStdout(), logger.GetStderr())
		}
	}
}

func Panic(arg0 interface{}, args ...interface{}) {

	if Verbosity >= PANIC {
		if logger == nil {
			showPreInit()
			output("panic  ", prepareMessage(arg0, args...), 2, "", nil, os.Stdout, os.Stderr)
		} else {
			output("panic  ", prepareMessage(arg0, args...), 2, logger.GetTimestampFormat(), logger.GetTemplate(), logger.GetStdout(), logger.GetStderr())
		}
	}
	panic("...")
}

func Fatal(arg0 interface{}, args ...interface{}) {

	if Verbosity >= FATAL {
		if logger == nil {
			showPreInit()
			output("fatal  ", prepareMessage(arg0, args...), 2, "", nil, os.Stdout, os.Stderr)
		} else {
			output("fatal  ", prepareMessage(arg0, args...), 2, logger.GetTimestampFormat(), logger.GetTemplate(), logger.GetStdout(), logger.GetStderr())
		}
	}
	os.Exit(1)
}

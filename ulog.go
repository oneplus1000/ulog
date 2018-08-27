package ulog

import (
	"bytes"
	"fmt"
	"log"
)

//LogLevelErr error
const LogLevelErr = 256

//LogLevelWarn warning
const LogLevelWarn = 128

//LogLevelInfo infomation
const LogLevelInfo = 64

//LogLevelDebug debug
const LogLevelDebug = 32

var writeLog = func(loglevel int, msg string) {
	log.Print(msg) //default
}

//SetWriteLog set writeLog function
func SetWriteLog(write func(loglevel int, msg string)) {
	writeLog = write
}

func logLevelDesc(loglevel int) string {
	if loglevel == LogLevelErr {
		return "error"
	} else if loglevel == LogLevelWarn {
		return "warn"
	} else if loglevel == LogLevelInfo {
		return "info"
	} else if loglevel == LogLevelDebug {
		return "debug"
	}
	return ""
}

//Errf error log
func Errf(err error, format string, a ...interface{}) {
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "%s: ", logLevelDesc(LogLevelErr))
	fmt.Fprintf(&buff, format, a...)
	fmt.Fprintf(&buff, "\n%+v", err)
	writeLog(LogLevelErr, buff.String())
}

//Err error log
func Err(err error, msg string) {
	Errf(err, msg)
}

//Info infomation log
func Info(msg string) {
	Infof(msg)
}

//Infof infomation log
func Infof(format string, a ...interface{}) {
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "%s: ", logLevelDesc(LogLevelInfo))
	fmt.Fprintf(&buff, format, a...)
	writeLog(LogLevelInfo, buff.String())
}

//Debug debug log
func Debug(msg string) {
	Debugf(msg)
}

//Debugf debug log
func Debugf(format string, a ...interface{}) {
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "%s: ", logLevelDesc(LogLevelDebug))
	fmt.Fprintf(&buff, format, a...)
	writeLog(LogLevelDebug, buff.String())
}

//Warn warning log
func Warn(msg string) {
	Warnf(msg)
}

//Warnf warning log
func Warnf(format string, a ...interface{}) {
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "%s: ", logLevelDesc(LogLevelWarn))
	fmt.Fprintf(&buff, format, a...)
	writeLog(LogLevelWarn, buff.String())
}

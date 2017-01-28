package util

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var GinLog *log.Logger
var RecLog *log.Logger
var AnyLog *log.Logger
var ContestGLog *log.Logger

func InitLogs(c Config) {
	var logLevel = log.DebugLevel.String()
	logDirPath := c.App().LogDirPath
	if len(logDirPath) == 0 {
		var lg = log.New()
		lg.Out = os.Stdout
		lg.Level = log.DebugLevel

		GinLog = lg
		RecLog = lg
		AnyLog = lg
		ContestGLog = lg
	} else {
		GinLog = newFileLog(logDirPath, logLevel, "gin.log")
		RecLog = newFileLog(logDirPath, logLevel, "rec.log")
		AnyLog = newFileLog(logDirPath, logLevel, "any.log")
		ContestGLog = newFileLog(logDirPath, logLevel, "contestG.log")
	}
}

func newFileLog(logDir, logLevel, logName string) *log.Logger {
	fileLog := &lumberjack.Logger{
		Filename:   logDir + "/" + logName,
		MaxSize:    5, // megabytes
		MaxBackups: 10,
		MaxAge:     28, //days
	}

	var lg = log.New()
	lg.Out = fileLog
	level, err := log.ParseLevel(logLevel)
	if err == nil {
		lg.Level = level
	}

	return lg
}

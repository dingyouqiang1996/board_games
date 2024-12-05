package logs

import (
    "common/config"
    "os"
    "time"

    "github.com/charmbracelet/log"
)

var logger *log.Logger

func InitLog(appName string) {
    logger = log.NewWithOptions(os.Stderr, log.Options{
        Prefix:          appName,
        ReportTimestamp: true,
        TimeFormat:      time.DateTime,
        Level:           log.InfoLevel,
    })
	
    if config.Conf.Log.Level == "DEBUG" {
        logger.SetLevel(log.DebugLevel)
    }
}

func Fatal(format string, values ...any) {
    if len(values) == 0 {
        logger.Fatal(format)
    } else {
        logger.Fatal(format, values...)
    }
}

func Info(format string, values ...any) {
	if len(values) == 0 {
        logger.Fatal(format)
    } else {
        logger.Fatalf(format, values...)
    }
}

func Debug(format string, values ...any) {
    if len(values) == 0 {
        logger.Debug(format)
    } else {
        logger.Debug(format, values...)
    }
}

func Error(format string, values ...any) {
    if len(values) == 0 {
        logger.Error(format)
    } else {
        logger.Error(format, values...)
    }
}
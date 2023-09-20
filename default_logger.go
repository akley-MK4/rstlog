package rstlog

import (
	"errors"
	"sync/atomic"
)

var (
	defaultLogger  = &Logger{}
	initializeDone uint32
)

func InitializeDefaultLogger(levelDesc string, outPrefix string, callDepth int) error {
	lv, exist := FindLogLevelByDesc(levelDesc)
	if !exist {
		return errors.New("invalid log level")
	}

	loggerInst, newErr := NewLogger(lv, callDepth, outPrefix, "", "")
	if newErr != nil {
		return newErr
	}

	if !atomic.CompareAndSwapUint32(&initializeDone, 0, 1) {
		return errors.New("default logger instance cannot be initialized repeatedly")
	}
	defaultLogger = loggerInst
	return nil
}

func GetDefaultLogger() *Logger {
	return defaultLogger
}

func SetDefaultLogLevel(logLevelDesc string) {
	lv, exist := FindLogLevelByDesc(logLevelDesc)
	if !exist {
		lv = LevelInfo
	}
	defaultLogger.SetLevel(lv)
}

func GetDefaultLogLevel() LogLevel {
	return defaultLogger.GetLogLevel()
}

func All(v ...interface{}) {
	if LevelALL < defaultLogger.level {
		return
	}
	defaultLogger.All(v...)
}

func AllF(format string, v ...interface{}) {
	if LevelALL < defaultLogger.level {
		return
	}
	defaultLogger.AllF(format, v...)
}

func Debug(v ...interface{}) {
	if LevelDebug < defaultLogger.level {
		return
	}
	defaultLogger.Debug(v...)
}

func DebugF(format string, v ...interface{}) {
	if LevelDebug < defaultLogger.level {
		return
	}
	defaultLogger.DebugF(format, v...)
}

func Info(v ...interface{}) {
	if LevelInfo < defaultLogger.level {
		return
	}
	defaultLogger.Info(v...)
}

func InfoF(format string, v ...interface{}) {
	if LevelInfo < defaultLogger.level {
		return
	}
	defaultLogger.InfoF(format, v...)
}

func Warning(v ...interface{}) {
	if LevelWarning < defaultLogger.level {
		return
	}
	defaultLogger.Warning(v...)
}

func WarningF(format string, v ...interface{}) {
	if LevelWarning < defaultLogger.level {
		return
	}
	defaultLogger.WarningF(format, v...)
}

func Error(v ...interface{}) {
	if LevelError < defaultLogger.level {
		return
	}
	defaultLogger.Error(v...)
}

func ErrorF(format string, v ...interface{}) {
	if LevelError < defaultLogger.level {
		return
	}
	defaultLogger.ErrorF(format, v...)
}

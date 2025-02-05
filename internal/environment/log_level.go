package environment

import (
	"go.uber.org/zap"
)

type LogLevel string

const (
	DebugLogLevel LogLevel = "DEBUG"
	InfoLogLevel  LogLevel = "INFO"
	WarnLogLevel  LogLevel = "WARN"
	ErrorLogLevel LogLevel = "ERROR"
)

func (logLevel LogLevel) String() string {
	return string(logLevel)
}

func (logLevel LogLevel) ZapLevel() zap.AtomicLevel {
	switch logLevel {
	case InfoLogLevel:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case WarnLogLevel:
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case ErrorLogLevel:
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	}
}

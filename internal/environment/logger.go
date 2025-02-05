package environment

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func createLogger(group string, level zap.AtomicLevel) *zap.Logger {
	currentProfile := Profile()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoding := "json"

	if currentProfile == "dev" {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoding = "console"
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	config := zap.Config{
		Level:             level,
		Development:       currentProfile == "dev",
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          encoding,
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		InitialFields: map[string]interface{}{
			"group": group,
		},
	}

	return zap.Must(config.Build())
}

func GetApplicationLogger() *zap.Logger {
	return createLogger("application", env.Log.Groups.Application.ZapLevel())
}

func GetPersistenceLogger() *zap.Logger {
	return createLogger("persistence", env.Log.Groups.Application.ZapLevel())
}

func GetPresentationLogger() *zap.Logger {
	return createLogger("presentation", env.Log.Groups.Application.ZapLevel())
}

func GetSecurityLogger() *zap.Logger {
	return createLogger("security", env.Log.Groups.Security.ZapLevel())
}

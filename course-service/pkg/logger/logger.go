package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.TimeKey = "timestamp"
	// encodeConfig.StacktraceKey = ""
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encodeConfig

	Log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	Log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	Log.Debug(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	Log.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	Log.Error(message, fields...)
}

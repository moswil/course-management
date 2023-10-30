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

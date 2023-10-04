package main

import (
	"context"
	"log"
	"os"

	"github.com/moswil/course-management/pkg/logger"
	"github.com/moswil/course-management/pkg/telemetry"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type tracingConfig struct {
	ServiceName string `json:"service_name"`
	ExporterURL string `json:"exporter_url"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
}

func initializeTracing() trace.Tracer {
	// let's unmarshal the tracing config
	var config tracingConfig
	err := viper.Unmarshal(&config)
	if err != nil {
		logger.Error("error marshalling: " + err.Error())
	}
	log.Printf("configs :%v\n", config)

	// tracing (init tracer)
	tracingProvider, err := telemetry.NewProvider(
		context.Background(),
		viper.GetString("tracing.service_name"),
		viper.GetString("tracing.exporter_url"),
		viper.GetString("app.version"),
		viper.GetString("app.environment"),
		viper.GetBool("tracing.secure_mode"),
	)
	if err != nil {
		logger.Error("error creating new provider: " + err.Error())
		os.Exit(1)
	}
	// registering tracing provider as global provider
	stopTracingProvider, err := tracingProvider.RegisterAsGlobal()
	if err != nil {
		logger.Error("error creating new provider: " + err.Error())
		os.Exit(1)
	}
	defer func() {
		if err := stopTracingProvider(context.TODO()); err != nil {
			logger.Info(err.Error())
			os.Exit(1)
		}
	}()
	tracer := otel.Tracer("")

	return tracer
}

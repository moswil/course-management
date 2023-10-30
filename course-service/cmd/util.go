package main

import (
	"context"
	"os"

	"github.com/moswil/course-management/pkg/telemetry"
	"github.com/spf13/viper"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
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
	ctx := context.Background()
	log := otelzap.L()
	// let's unmarshal the tracing config
	var config tracingConfig
	err := viper.Unmarshal(&config)
	if err != nil {
		log.ErrorContext(ctx, "error marshalling: "+err.Error())
	}

	// tracing (init tracer)
	tracingProvider, err := telemetry.NewProvider(
		context.Background(),
		viper.GetString("telemetry.service_name"),
		viper.GetString("telemetry.exporter_url"),
		viper.GetString("app.version"),
		viper.GetString("app.environment"),
		viper.GetBool("telemetry.secure_mode"),
	)
	if err != nil {
		log.ErrorContext(ctx, "error creating new provider: "+err.Error())
		os.Exit(1)
	}
	// registering tracing provider as global provider
	stopTracingProvider, err := tracingProvider.RegisterAsGlobal()
	if err != nil {
		log.ErrorContext(ctx, "error creating new provider: "+err.Error())
		os.Exit(1)
	}
	defer func() {
		if err := stopTracingProvider(context.TODO()); err != nil {
			log.InfoContext(ctx, err.Error())
			os.Exit(1)
		}
	}()
	tracer := otel.Tracer("")

	return tracer
}

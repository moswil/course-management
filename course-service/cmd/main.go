package main

import (
	"context"
	"fmt"
	"os"
	"path"

	grpc "github.com/moswil/course-management/internal/app/grpc/server"
	grpcSvc "github.com/moswil/course-management/internal/app/grpc/service"
	rest "github.com/moswil/course-management/internal/app/rest/server"
	"github.com/moswil/course-management/internal/core/service"
	"github.com/moswil/course-management/internal/core/util/constants"
	"github.com/moswil/course-management/internal/infra/messaging"
	sqlUtil "github.com/moswil/course-management/internal/infra/repository/sql/util"
	"github.com/moswil/course-management/internal/infra/util"
	"github.com/moswil/course-management/pkg/logger"
	"github.com/spf13/viper"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap/zapcore"
)

var tracer trace.Tracer

func init() {
	// Configure Viper
	currDir, _ := os.Getwd()
	util.SetupConfig(path.Join(currDir, "./conf"), "app", "yaml")
}

func init() {
	tracer = initializeTracing()
}

func main() {
	// get logger for telemetry
	var appLogger *otelzap.Logger
	if getEnvironment() == constants.ENV_DEVELOPMENT {
		appLogger = otelzap.New(
			logger.Log,
			otelzap.WithTraceIDField(true),           // for development
			otelzap.WithMinLevel(zapcore.DebugLevel), // for development
		)
	} else {
		appLogger = otelzap.New(
			logger.Log,
		)
	}

	// sugar := logger.Sugar()
	defer appLogger.Sync()

	undo := otelzap.ReplaceGlobals(appLogger)
	defer undo()

	// use logger
	log := otelzap.L()
	ctx := context.TODO()

	// Initialize repositories, services, and servers

	// repository
	courseRepository, err := sqlUtil.InitializeMySQLRepository(tracer)
	if err != nil {
		errMsg := fmt.Sprintf("error: %v occurred connecting to db", err.Error())
		// logger.Error(errMsg)
		log.ErrorContext(ctx, errMsg)
	}

	//  broker (kafka)
	// eventBroker := messaging.NewKafkaEventPublisher()

	eventBrokerProtobuf := messaging.NewKafkaProtobufEventPublisher()

	// course service
	// courseService := service.NewCourseService(courseRepository, eventBroker)
	courseService := service.NewCourseService(courseRepository, eventBrokerProtobuf)

	// Start gRPC Server as a goroutine
	go func() {
		grpcPort := 50052
		grpcServer := grpc.NewGRPCServer(grpcPort, *grpcSvc.NewCourseService(courseService))
		if err := grpcServer.Start(); err != nil {
			log.FatalContext(ctx, "Failed to start gRPC server: "+err.Error())
		}
	}()

	// Start the REST API server
	go func() {
		restServer := rest.NewRestServer(courseService, tracer)
		if err := restServer.Start(":8080"); err != nil {
			log.FatalContext(ctx, "Failed to start REST server: "+err.Error())
		}
	}()

	// optionally wait for the go routines
	select {}
}

func getEnvironment() string {
	return viper.GetString("app.environment")
}

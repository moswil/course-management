// provider.go initializes the tracing provider.

package telemetry

import (
	"context"
	"time"

	"github.com/moswil/course-management/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc/credentials"
)

type Provider struct {
	serviceName string
	exporterURL string
	provider    *trace.TracerProvider
}

func NewProvider(ctx context.Context, serviceName, exporterURL, appVersion, environment string, secureMode bool) (*Provider, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// setup exporter
	// we create an OTLP exporter with a tracing backend endpoint
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))

	if !secureMode {
		secureOption = otlptracegrpc.WithInsecure()
	}

	e, err := otlptrace.New(ctx, otlptracegrpc.NewClient(
		otlptracegrpc.WithEndpoint(exporterURL),
		secureOption,
	))
	if err != nil {
		logger.Error("error creating exporter: " + err.Error())
		return nil, err
	}

	// setup resource, which describes an entity about identifying information and metadata is exposed
	r := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
		semconv.ServiceVersionKey.String(appVersion),
		attribute.String("environment", environment),
	)

	// setup sampler (NOT GOOD FOR PRODUCTION)
	// AlwaysSample returns a Sampler that samples every trace, for testing purpose we want to sample every trace
	// but in the production system with high number of requests it can be expensive
	s := trace.AlwaysSample()

	traceProvider := trace.NewTracerProvider(
		trace.WithSampler(s),
		trace.WithBatcher(e),
		trace.WithResource(r),
	)

	return &Provider{
		serviceName: serviceName,
		exporterURL: exporterURL,
		provider:    traceProvider,
	}, nil
}

func (p *Provider) RegisterAsGlobal() (func(ctx context.Context) error, error) {
	// set global provider
	otel.SetTracerProvider(p.provider)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return p.provider.Shutdown, nil
}

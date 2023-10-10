package tracing

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"log"
)

func SetupTracer(ctx context.Context, cfg *config.Config) *trace.TracerProvider {
	var exp *otlptrace.Exporter
	var err error

	exp, err = otlptracehttp.New(
		ctx,
		otlptracehttp.WithEndpoint(cfg.JaegerCollectorHost),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to create otpl exporters: %s", err.Error())
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("app-auth"),
			attribute.String("environment", cfg.Env),
		)),
	)

	return traceProvider
}

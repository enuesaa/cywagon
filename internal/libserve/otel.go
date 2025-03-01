package libserve

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"go.opentelemetry.io/otel/trace"
)


type OtelMiddleware struct {
	Next Middleware
}

func (m *OtelMiddleware) Handle(site Site, req *http.Request) (*http.Response, error) {
	if tracer == nil {
		return m.Next.Handle(site, req)
	}

	_, span := tracer.Start(context.Background(), "handle")
	defer span.End()

	return m.Next.Handle(site, req)
}

var tracer trace.Tracer

func setupTracer(ctx context.Context) error {
	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint("localhost:4318"),
		otlptracehttp.WithInsecure(),
	)
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return err
	}

	resources, err := newTraceResoure(ctx)
	if err != nil {
		return err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
	)
	otel.SetTracerProvider(tp)

	tracer = otel.Tracer("localhost:3000")

	return nil
}

func newTraceResoure(ctx context.Context) (*resource.Resource, error) {
	return resource.New(
		ctx,
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceName("aa"),
		),
	)
}
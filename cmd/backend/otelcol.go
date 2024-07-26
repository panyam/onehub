package main

import (
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func newOtelTraceProvider[C any](o *OTELSetup[C]) (trace.TracerProvider, ShutdownFunc, error) {
	// Set up a trace exporter
	traceExporter, err := otlptracegrpc.New(o.ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// Register the trace exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(o.Resource),
		sdktrace.WithSpanProcessor(bsp),
	)

	return tracerProvider, tracerProvider.Shutdown, nil
}

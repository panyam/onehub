package main

import (
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	otelmetric "go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type CollectorSetup struct {
	ClientConn *grpc.ClientConn
}

func NewOTELSetupWithCollector(conn *grpc.ClientConn) (out *OTELSetup[CollectorSetup]) {
	out = &OTELSetup[CollectorSetup]{
		SetupTracerProvider: newOtelTraceProvider,
		SetupMeterProvider:  newOtelMeterProvider,
		SetupLoggerProvider: newStdoutLoggerProvider[CollectorSetup],
	}
	out.Context.ClientConn = conn
	return
}

func newOtelTraceProvider(o *OTELSetup[CollectorSetup]) (trace.TracerProvider, ShutdownFunc, error) {
	// Set up a trace exporter
	traceExporter, err := otlptracegrpc.New(o.ctx, otlptracegrpc.WithGRPCConn(o.Context.ClientConn))
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

func newOtelMeterProvider(o *OTELSetup[CollectorSetup]) (otelmetric.MeterProvider, ShutdownFunc, error) {
	metricExporter, err := otlpmetricgrpc.New(o.ctx, otlpmetricgrpc.WithGRPCConn(o.Context.ClientConn))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create metrics exporter: %w", err)
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter)),
		sdkmetric.WithResource(o.Resource),
	)
	return meterProvider, meterProvider.Shutdown, nil
}

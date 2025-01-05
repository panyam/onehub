package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	cmdutils "github.com/panyam/onehub/cmd/utils"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	otelmetric "go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CollectorSetup struct {
	ClientConn *grpc.ClientConn
}

func SetupOtel(collectorAddrEnv string, defaultCollectorAddr string) (context.Context, *OTELSetup[CollectorSetup], context.CancelFunc, error) {
	if collectorAddrEnv == "" {
		collectorAddrEnv = "OTEL_COLLECTOR_ADDR"
	}

	if defaultCollectorAddr == "" {
		defaultCollectorAddr = "otel-collector:4317"
	}
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	collectorAddr := cmdutils.GetEnvOrDefault("OTEL_COLLECTOR_ADDR", "otel-collector:4317")
	conn, err := grpc.NewClient(collectorAddr,
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("failed to create gRPC connection to collector: %w", err)
		return nil, nil, nil, err
	}
	setup := NewOTELSetupWithCollector(conn)
	err = setup.Setup(ctx)
	if err != nil {
		log.Println("error setting up otel: ", err)
	}
	return ctx, setup, stop, err
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

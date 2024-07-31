package main

import (
	"time"

	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	otelmetric "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func NewOTELSetupWithStdout() (out *OTELSetup[any]) {
	out = &OTELSetup[any]{
		SetupTracerProvider: newStdoutTraceProvider[any],
		SetupMeterProvider:  newStdoutMeterProvider,
		SetupLoggerProvider: newStdoutLoggerProvider[any],
	}
	return
}

func newStdoutMeterProvider(o *OTELSetup[any]) (otelmetric.MeterProvider, ShutdownFunc, error) {
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter,
			// Default is 1m. Set to 3s for demonstrative purposes.
			metric.WithInterval(3*time.Second))),
	)
	return meterProvider, meterProvider.Shutdown, nil
}

func newStdoutLoggerProvider[C any](o *OTELSetup[C]) (*log.LoggerProvider, ShutdownFunc, error) {
	logExporter, err := stdoutlog.New()
	if err != nil {
		return nil, nil, err
	}

	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)
	return loggerProvider, loggerProvider.Shutdown, nil
}

func newStdoutTraceProvider[C any](o *OTELSetup[C]) (trace.TracerProvider, ShutdownFunc, error) {
	traceExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, nil, err
	}

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter,
			// Default is 5s. Set to 1s for demonstrative purposes.
			sdktrace.WithBatchTimeout(time.Second)),
		sdktrace.WithResource(o.Resource),
	)
	return traceProvider, traceProvider.Shutdown, nil
}

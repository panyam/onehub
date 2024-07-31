package main

import (
	"context"
	"errors"
	sl "log"

	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
	otelmetric "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type ShutdownFunc = func(context.Context) error

type OTELSetup[C any] struct {
	ctx                 context.Context
	shutdownFuncs       []ShutdownFunc
	Resource            *resource.Resource
	Context             C
	SetupPropagator     func(o *OTELSetup[C])
	SetupTracerProvider func(o *OTELSetup[C]) (trace.TracerProvider, ShutdownFunc, error)
	SetupMeterProvider  func(o *OTELSetup[C]) (otelmetric.MeterProvider, ShutdownFunc, error)
	SetupLogger         func(o *OTELSetup[C]) (logr.Logger, ShutdownFunc, error)
	SetupLoggerProvider func(o *OTELSetup[C]) (*log.LoggerProvider, ShutdownFunc, error)
}

func (o *OTELSetup[C]) Shutdown(ctx context.Context) error {
	var err error
	for _, fn := range o.shutdownFuncs {
		err = errors.Join(err, fn(ctx))
	}
	o.shutdownFuncs = nil
	return err
}

func (o *OTELSetup[C]) HandleError(inErr error) error {
	return errors.Join(inErr, o.Shutdown(o.ctx))
}

func (o *OTELSetup[C]) Setup(ctx context.Context) (err error) {
	// Ensure default SDK resources and the required service name are set.
	o.ctx = ctx
	o.Resource, err = resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("onehub")),
	)

	if err != nil {
		sl.Println("err: ", err)
		return
	}
	if o.SetupPropagator != nil {
		o.SetupPropagator(o)
	} else {
		otel.SetTextMapPropagator(newPropagator())
	}

	if o.SetupTracerProvider != nil {
		tp, sf, err := o.SetupTracerProvider(o)
		if err != nil {
			sl.Println("TraceProvider Error: ", err)
			err = o.HandleError(err)
			return err
		} else {
			otel.SetTracerProvider(tp)
			o.shutdownFuncs = append(o.shutdownFuncs, sf)
		}
	} else {
		sl.Println("No tracer provider")
	}

	if o.SetupMeterProvider != nil {
		mp, sf, err := o.SetupMeterProvider(o)
		if err != nil {
			sl.Println("MeterProvider Error: ", err)
			err = o.HandleError(err)
			return err
		} else {
			otel.SetMeterProvider(mp)
			o.shutdownFuncs = append(o.shutdownFuncs, sf)
		}
	} else {
		sl.Println("No meter provider")
	}

	if o.SetupLoggerProvider != nil {
		lp, sf, err := o.SetupLoggerProvider(o)
		if err != nil {
			sl.Println("LoggerProvider Error: ", err)
			err = o.HandleError(err)
			return err
		} else {
			global.SetLoggerProvider(lp)
			o.shutdownFuncs = append(o.shutdownFuncs, sf)
		}
	} else {
		sl.Println("No logger provider")
	}
	return nil
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

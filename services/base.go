package services

import (
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

const name = "github.com/panyam/onehub"

var (
	tracer     = otel.Tracer(name)
	meter      = otel.Meter(name)
	logger     = otelslog.NewLogger(name)
	userCnt    metric.Int64Counter
	topicCnt   metric.Int64Counter
	messageCnt metric.Int64Counter
)

func init() {
	var err error
	userCnt, err = meter.Int64Counter("onehub.users",
		metric.WithDescription("The number of users created in this system"),
		metric.WithUnit("{user}"))
	if err != nil {
		panic(err)
	}
	topicCnt, err = meter.Int64Counter("onehub.topics",
		metric.WithDescription("The number of topics created in this system"),
		metric.WithUnit("{topic}"))
	if err != nil {
		panic(err)
	}
	messageCnt, err = meter.Int64Counter("onehub.messages",
		metric.WithDescription("The number of messages created in this system"),
		metric.WithUnit("{message}"))
	if err != nil {
		panic(err)
	}
}

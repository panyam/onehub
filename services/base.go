package services

import (
	"github.com/panyam/onehub/obs"
	"go.opentelemetry.io/otel/metric"
)

var (
	userCnt    metric.Int64Counter
	topicCnt   metric.Int64Counter
	messageCnt metric.Int64Counter
)

func init() {
	var err error
	userCnt, err = obs.Meter.Int64Counter("onehub.users",
		metric.WithDescription("The number of users created in this system"),
		metric.WithUnit("{user}"))
	if err != nil {
		panic(err)
	}
	topicCnt, err = obs.Meter.Int64Counter("onehub.topics",
		metric.WithDescription("The number of topics created in this system"),
		metric.WithUnit("{topic}"))
	if err != nil {
		panic(err)
	}
	messageCnt, err = obs.Meter.Int64Counter("onehub.messages",
		metric.WithDescription("The number of messages created in this system"),
		metric.WithUnit("{message}"))
	if err != nil {
		panic(err)
	}
}

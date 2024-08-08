package obs

import (
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
)

const name = "github.com/panyam/onehub"

var (
	Tracer = otel.Tracer(name)
	Meter  = otel.Meter(name)
	Logger = otelslog.NewLogger(name)
)

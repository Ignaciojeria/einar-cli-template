package business

import (
	"go.opentelemetry.io/otel"
)

var Tracer = otel.Tracer("business")

package business

import (
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("business")

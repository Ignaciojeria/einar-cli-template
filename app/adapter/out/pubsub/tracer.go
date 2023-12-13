package pubsub

import (
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("publisher")

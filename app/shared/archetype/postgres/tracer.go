package postgres

import (
	"go.opentelemetry.io/otel"
)

var Tracer = otel.Tracer("postgresql")

package postgresql

import (
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("postgresql")

package client

import (
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("resty")

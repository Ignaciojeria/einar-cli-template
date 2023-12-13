package resty

import "go.opentelemetry.io/otel"

var Tracer = otel.Tracer("resty")

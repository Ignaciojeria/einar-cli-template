package topic

import "go.opentelemetry.io/otel"

var Tracer = otel.Tracer("pubsub-publisher")

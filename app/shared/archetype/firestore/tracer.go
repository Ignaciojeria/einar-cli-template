package firestore

import "go.opentelemetry.io/otel"

var Tracer = otel.Tracer("firestore")

package firestore

import (
	einar "archetype/app/shared/archetype/firestore"
	"context"

	"cloud.google.com/go/firestore"
	"go.opentelemetry.io/otel/trace"
)

var ArchetypeRepository = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) (err error) {
	_, span := einar.Tracer.Start(ctx,
		"ArchetypeRepository",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	var _ *firestore.CollectionRef = einar.Collection("INSERT_YOUR_COLLECTION_CONSTANT_HERE")
	return nil
}

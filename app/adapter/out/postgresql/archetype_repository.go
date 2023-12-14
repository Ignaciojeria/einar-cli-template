package postgresql

import (
	einar "archetype/app/shared/archetype/postgres"
	"context"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

var ArchetypeRepository = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) error {
	_, span := einar.Tracer.Start(ctx,
		"ArchetypeRepository",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	var _ *gorm.DB = einar.DB
	//PUT YOUR POSTGRESL OPERATION USING EINAR HERE :
	//....einar.DB....
	return nil
}

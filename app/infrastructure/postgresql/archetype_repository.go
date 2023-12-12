package postgresql

import (
	"archetype/app/domain/ports/out"
	einar "archetype/app/shared/archetype/postgres"
	"archetype/app/shared/ddlog"
	"context"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gorm.io/gorm"
)

var ArchetypeRepository out.ArchetypeOutBoundPort = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) (err error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "ArchetypeRepository", tracer.AnalyticsRate(1))
	defer func() {
		if err != nil {
			span.SetTag("error.msg", err.Error())
			ddlog.
				Err(span.Context()).
				Err(err).
				Msg("ArchetypeRepository" + "Error")
			span.Finish(tracer.WithError(err))
		} else {
			span.Finish()
		}
	}()

	var _ *gorm.DB = einar.DB
	//PUT YOUR POSTGRESL OPERATION USING EINAR HERE :
	//....einar.DB....
	return nil
}

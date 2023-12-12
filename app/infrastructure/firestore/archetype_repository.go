package firestore

import (
	"archetype/app/domain/ports/out"
	"archetype/app/shared/archetype/exception"
	einar "archetype/app/shared/archetype/firestore"
	"archetype/app/shared/ddlog"
	"context"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
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

	archetypeCollection := einar.Collection("INSERT_YOUR_COLLECTION_PATH_CONSTANT_HERE")
	if archetypeCollection == nil {
		return exception.CollectionIsNotPresentError{}
	}
	//PUT YOUR FIRESTORE OPERATION HERE :
	//....archetypeCollection.Doc()
	//....archetypeCollection.Add()
	//....
	return nil
}

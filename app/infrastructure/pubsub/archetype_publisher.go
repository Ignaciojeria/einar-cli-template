package pubsub

import (
	"archetype/app/domain/ports/out"
	einar "archetype/app/shared/archetype/pubsub"
	"archetype/app/shared/archetype/pubsub/publisher"
	"archetype/app/shared/ddlog"
	"context"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var ArchetypePublisher out.ArchetypeOutBoundPort = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) (err error) {

	span, ctx := tracer.StartSpanFromContext(ctx, "ArchetypePublisher", tracer.AnalyticsRate(1))
	defer func() {
		if err != nil {
			span.SetTag("error.msg", err.Error())
			ddlog.
				Err(span.Context()).
				Err(err).
				Msg("ArchetypePublisher" + "Error")
			span.Finish(tracer.WithError(err))
		} else {
			span.Finish()
		}
	}()

	if err := publisher.PublishAtTopic(ctx, REPLACE_BY_YOUR_DOMAIN, einar.Topic("INSERT YOUR TOPIC NAME HERE"), map[string]string{
		"customAttribute1": "REPLACE_BY_YOUR CUSTOM ATTRIBUTES",
		"customAttribute2": "....",
	}); err != nil {
		return err
	}
	return nil
}

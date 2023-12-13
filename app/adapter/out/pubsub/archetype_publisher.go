package pubsub

import (
	"archetype/app/exception"
	"archetype/app/shared/archetype/pubsub/topic"
	"archetype/app/shared/archetype/slog"
	"archetype/app/shared/constants"
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var ArchetypePublisher = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) (err error) {
	topicName := "INSERT YOUR TOPIC NAME HERE"

	_, span := topic.Tracer.Start(ctx, "ArchetypePublisher",
		trace.WithAttributes(attribute.String(constants.TopicName, topicName)),
	)
	defer span.End()

	bytes, err := json.Marshal(REPLACE_BY_YOUR_DOMAIN)
	if err != nil {
		return err
	}

	message := &pubsub.Message{
		Attributes: map[string]string{
			"customAttribute1": "attr1",
			"customAttribute2": "attr2",
		},
		Data: bytes,
	}
	result := topic.Get(topicName).Publish(ctx, message)
	// Get the server-generated message ID.
	messageID, err := result.Get(ctx)

	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)

		slog.
			SpanLogger(span).
			Error(exception.PUBSUB_BROKER_ERROR.Error(),
				constants.Error, err.Error())

		return exception.PUBSUB_BROKER_ERROR
	}

	// Successful publishing
	fmt.Println("Message published with ID:", messageID)
	return nil
}

package publisher

import (
	"archetype/app/shared/utils"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func PublishAtTopic(ctx context.Context, body interface{}, t *pubsub.Topic, attr map[string]string) error {

	b, err := json.Marshal(body)

	if err != nil {
		return err
	}

	attributes := make(map[string]string)
	attributes["eventId"] = uuid.NewString()
	attributes["timestamp"] = fmt.Sprintf("%d", int32(time.Now().Unix()))
	attributes["datetime"] = time.Now().UTC().Format(time.RFC3339)

	for k, v := range attr {
		attributes[k] = v
	}

	attr = attributes

	msg := &pubsub.Message{
		Data:       utils.CompressBytes(b),
		Attributes: attr,
	}

	if _, err = Publish(ctx, t, msg).Get(ctx); err != nil {

		log.Error().
			Err(err).
			Str("topic.name", t.String()).
			Interface("msg.body", body).
			Interface("msg.attributes", attributes).
			Msg("Error was trying a message")

		return err
	}

	log.Info().
		Str("topic.name", t.String()).
		Interface("msg.body", body).
		Interface("msg.attributes", attributes).
		Msg("A message was published")

	return nil

}

package pubsub

import (
	einar "archetype/app/shared/archetype/pubsub"
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
)

var ArchetypePublisher = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) (err error) {

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

	result := einar.Topic("INSERT YOUR TOPIC NAME HERE").Publish(ctx, message)

	// Get the server-generated message ID.
	messageID, err := result.Get(ctx)
	if err != nil {
		return err
	}

	// Successful publishing
	fmt.Println("Message published with ID:", messageID)
	return nil
}

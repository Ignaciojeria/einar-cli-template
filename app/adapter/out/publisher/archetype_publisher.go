package publisher

import (
	"archetype/app/domain/ports/out"
	einar "archetype/app/shared/archetype/pubsub"
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
)

var ArchetypePublisher out.ArchetypeOutBoundPort = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) error {
	var p *pubsub.Topic = einar.Client.Topic("INSERT_YOUR_TOPIC_NAME_HERE")

	bytes, err := json.Marshal(REPLACE_BY_YOUR_DOMAIN)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	message := &pubsub.Message{
		Attributes: map[string]string{
			"customAttribute1": "attr1",
			"customAttribute2": "attr2",
		},
		Data: bytes,
	}

	result := p.Publish(ctx, message)

	// Get the server-generated message ID.
	messageID, err := result.Get(ctx)
	if err != nil {
		// Handle the error
		fmt.Println("Error occurred while publishing the result:", err.Error())
		// Perform any necessary error handling actions
		return err
	}

	// Successful publishing
	fmt.Println("Message published with ID:", messageID)
	return nil
}

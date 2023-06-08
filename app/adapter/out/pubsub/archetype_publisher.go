package pubsub

import (
	"archetype/app/domain/ports/out"
	"archetype/app/shared/archetype/container"
	einar "archetype/app/shared/archetype/pubsub"
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
)

var ArchetypeTopic *pubsub.Topic

func init() {
	container.InjectComponent(func() error {
		ArchetypeTopic = einar.Client.Topic("INSERT YOUR TOPIC NAME HERE")
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
		Paralel:      false,
	})
}

var ArchetypePublisher out.ArchetypeOutBoundPort = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) error {

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

	result := ArchetypeTopic.Publish(ctx, message)

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

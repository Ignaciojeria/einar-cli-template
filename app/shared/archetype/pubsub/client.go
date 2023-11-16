package pubsub

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/archetype/slog"
	"archetype/app/shared/config"
	"archetype/app/shared/constants"
	"context"
	"sync"

	"cloud.google.com/go/pubsub"
)

var (
	Client    *pubsub.Client
	topicRefs sync.Map
)

func init() {
	config.Installations.EnablePubSub = true
	container.InjectInstallation(func() error {
		projectId := config.GOOGLE_PROJECT_ID.Get()
		c, err := pubsub.NewClient(context.Background(), projectId)
		if err != nil {
			slog.Logger.Error("error getting pubsub client", constants.ERROR, err.Error())
			return err
		}
		Client = c
		return nil
	})
}

// Topic fetches a *pubsub.Topic by name. If the Topic exists in the sync.Map, it's returned, otherwise a new one is created and stored in the map.
func Topic(topicName string) *pubsub.Topic {
	value, ok := topicRefs.Load(topicName)
	if ok {
		// If the topic reference was found, return it.
		return value.(*pubsub.Topic)
	}

	// If the topic reference was not found, create a new one.
	newTopicRef := Client.Topic(topicName)

	// Store the new topic reference in the map.
	topicRefs.Store(topicName, newTopicRef)

	// Return the new topic reference.
	return newTopicRef
}

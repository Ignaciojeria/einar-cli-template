package pubsub

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/config"
	"archetype/app/shared/utils"
	"context"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

var (
	Client    *pubsub.Client
	topicRefs sync.Map
)

func init() {
	config.Installations.EnablePubSub = true
	container.InjectInstallation(func() error {
		projectId := config.GOOGLE_PROJECT_ID.Get()
		creds, err := utils.DecodeBase64(config.GOOGLE_APPLICATION_CRETENTIALS_B64.Get())
		if err != nil {
			log.Error().Err(err).Msg("error decoding GOOGLE_APPLICATION_CRETENTIALS_B64")
			return err
		}
		c, err := pubsub.NewClient(context.Background(), projectId, option.WithCredentialsJSON(creds))
		if err != nil {
			log.Error().Err(err).Msg("error getting pubsub client")
			return err
		}
		Client = c
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
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

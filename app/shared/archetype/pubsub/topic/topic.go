package topic

import (
	einar "archetype/app/shared/archetype/pubsub"
	"sync"

	"cloud.google.com/go/pubsub"
)

var topicRefs sync.Map

// Topic fetches a *pubsub.Topic by name. If the Topic exists in the sync.Map, it's returned, otherwise a new one is created and stored in the map.
func Get(topicName string) *pubsub.Topic {
	value, ok := topicRefs.Load(topicName)
	if ok {
		// If the topic reference was found, return it.
		return value.(*pubsub.Topic)
	}
	// If the topic reference was not found, create a new one.
	newTopicRef := einar.Client().Topic(topicName)
	// Store the new topic reference in the map.
	topicRefs.Store(topicName, newTopicRef)

	// Return the new topic reference.
	return newTopicRef
}

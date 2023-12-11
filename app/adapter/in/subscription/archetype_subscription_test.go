package subscription

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
)

func TestArchetypeSubscriptionSuccedded(t *testing.T) {
	archetypeInboundAdapterConfig()
	ctx := context.Background()
	testData := `{"key": "value"}`
	msg := &pubsub.Message{
		Data: []byte(testData),
	}
	err := archetype_subscription(ctx, "fake-subscription-dont-change", msg)
	if err != nil {
		t.Errorf("archetype_subscription returned an error: %v", err)
	}
}

func TestArchetypeSubscriptionInvalidInput(t *testing.T) {
	ctx := context.Background()
	invalidTestData := `invalid json`
	msg := &pubsub.Message{
		Data: []byte(invalidTestData),
	}
	err := archetype_subscription(ctx, "fake-subscription-dont-change", msg)
	if err == nil {
		t.Errorf("archetype_subscription did not return an error for invalid input")
	}
}

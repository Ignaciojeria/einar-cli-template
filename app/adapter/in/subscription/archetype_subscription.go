package subscription

import (
	"context"
	"encoding/json"

	"archetype/app/exception"
	"archetype/app/shared/archetype/container"
	einar "archetype/app/shared/archetype/pubsub"
	"archetype/app/shared/archetype/pubsub/subscription"
	"archetype/app/shared/constants"

	"cloud.google.com/go/pubsub"
)

func init() {
	const subscriptionName = "INSERT YOUR SUBSCRIPTION NAME"

	processMessage := func(ctx context.Context, m *pubsub.Message) (err error) {
		var dataModel interface{}
		defer subscription.HandleMessageAcknowledgement(ctx, &subscription.HandleMessageAcknowledgementDetails{
			SubscriptionName: subscriptionName,
			Error:            err,
			Message:          m,
			ErrorsRequiringNack: []error{
				exception.INTERNAL_SERVER_ERROR,
			},
			CustomLogFields: map[string]interface{}{
				"dataModel": dataModel,
			},
		})
		if json.Unmarshal(m.Data, &dataModel) != nil {
			return err
		}
		return nil
	}

	container.InjectInboundAdapter(func() error {
		subRef := einar.Client.Subscription(subscriptionName)
		subRef.ReceiveSettings.MaxOutstandingMessages = 5
		settings := subRef.Receive
		go subscription.
			New(subscriptionName, processMessage, settings).
			WithPushHandler(constants.DefaultPushHandlerPrefix + subscriptionName).
			Start()
		return nil
	})
}

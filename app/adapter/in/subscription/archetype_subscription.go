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

var archetype_subscription = func(ctx context.Context, subscriptionName string, m *pubsub.Message) (err error) {
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
	if err := json.Unmarshal(m.Data, &dataModel); err != nil {
		return err
	}
	return nil
}

var init_archetype_subscription = func() error {
	var subscriptionName = "INSERT YOUR SUBSCRIPTION NAME"
	subRef := einar.Client().Subscription(subscriptionName)
	subRef.ReceiveSettings.MaxOutstandingMessages = 5
	settings := subRef.Receive
	go subscription.
		New(subscriptionName, archetype_subscription, settings).
		WithPushHandler(constants.DefaultPushHandlerPrefix + subscriptionName).
		Start()
	return nil
}

var _ = container.InjectInboundAdapter(init_archetype_subscription)

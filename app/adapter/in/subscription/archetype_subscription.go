package subscription

import (
	"context"
	"encoding/json"
	"net/http"

	"archetype/app/exception"
	"archetype/app/shared/archetype/container"
	einar "archetype/app/shared/archetype/pubsub"
	"archetype/app/shared/archetype/pubsub/subscription"
	"archetype/app/shared/constants"

	"cloud.google.com/go/pubsub"
)

var archetype_subscription = func(ctx context.Context, subscriptionName string, m *pubsub.Message) (statusCode int, err error) {
	var dataModel interface{}
	defer func() {
		subscription.HandleMessageAcknowledgement(ctx, &subscription.HandleMessageAcknowledgementDetails{
			MessageID:        m.ID,
			PublishTime:      m.PublishTime.String(),
			SubscriptionName: subscriptionName,
			Error:            err,
			Message:          m,
			ErrorsRequiringNack: []error{
				exception.INTERNAL_SERVER_ERROR,
				exception.EXTERNAL_SERVER_ERROR,
				exception.HTTP_NETWORK_ERROR,
				exception.PUBSUB_BROKER_ERROR,
			},
			CustomLogFields: map[string]interface{}{
				"dataModel": dataModel,
			},
		})
	}()
	if err := json.Unmarshal(m.Data, &dataModel); err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

var archetypeInboundAdapterConfig = func() error {
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

var _ = container.InjectInboundAdapter(archetypeInboundAdapterConfig)

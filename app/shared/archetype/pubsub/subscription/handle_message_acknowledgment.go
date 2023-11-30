package subscription

import (
	"archetype/app/shared/archetype/slog"
	"archetype/app/shared/constants"
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
)

type HandleMessageAcknowledgementDetails struct {
	SubscriptionName    string
	Error               error
	Message             *pubsub.Message
	ErrorsRequiringNack []error
}

func HandleMessageAcknowledgement(ctx context.Context, details *HandleMessageAcknowledgementDetails) {
	if details.Error != nil {
		slog.Logger.Error(
			details.SubscriptionName+"_exception",
			constants.SUBSCRIPTION_NAME, details.SubscriptionName,
			constants.ERROR, details.Error,
		)
		for _, err := range details.ErrorsRequiringNack {
			if errors.Is(details.Error, err) {
				return details.Message.Nack()
			}
		}
		return details.Message.Ack()
	}
	slog.Logger.Info(
		details.SubscriptionName+"_succedded",
		constants.SUBSCRIPTION_NAME, details.SubscriptionName,
		constants.ERROR, details.Error,
	)

	details.Message.Ack()
}

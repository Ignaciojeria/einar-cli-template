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
	CustomLogFields     map[string]interface{}
}

func HandleMessageAcknowledgement(ctx context.Context, details *HandleMessageAcknowledgementDetails) {
	if details.Error != nil {
		slog.Logger.Error(
			details.SubscriptionName+"_exception",
			subscription_name, details.SubscriptionName,
			constants.Fields, details.CustomLogFields,
			constants.Error, details.Error,
		)
		for _, err := range details.ErrorsRequiringNack {
			if errors.Is(details.Error, err) {
				details.Message.Nack()
				return
			}
		}
		details.Message.Ack()
		return
	}
	// Logica existente para manejar el mensaje exitoso
	slog.Logger.Info(
		details.SubscriptionName+"_succedded",
		subscription_name, details.SubscriptionName,
		constants.Fields, details.CustomLogFields,
	)
	details.Message.Ack()
}

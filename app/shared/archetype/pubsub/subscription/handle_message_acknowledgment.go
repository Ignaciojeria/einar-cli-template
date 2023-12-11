package subscription

import (
	"archetype/app/shared/archetype/slog"
	"archetype/app/shared/constants"
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type HandleMessageAcknowledgementDetails struct {
	SubscriptionName    string
	Error               error
	Message             *pubsub.Message
	ErrorsRequiringNack []error
	CustomLogFields     map[string]interface{}
}

func HandleMessageAcknowledgement(ctx context.Context, details *HandleMessageAcknowledgementDetails) {
	tracer := otel.Tracer("pubsub-tracer")
	ctx, span := tracer.Start(ctx, "HandleMessageAcknowledgement")
	defer span.End()

	span.SetAttributes(
		attribute.String("subscription.name", details.SubscriptionName),
		attribute.String("message.id", details.Message.ID),
		attribute.String("message.publishTime", details.Message.PublishTime.String()),
	)

	if details.Error != nil {
		span.RecordError(details.Error)
		span.SetStatus(codes.Error, details.Error.Error())
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
	slog.Logger.Info(
		details.SubscriptionName+"_succedded",
		subscription_name, details.SubscriptionName,
		constants.Fields, details.CustomLogFields,
	)
	details.Message.Ack()
}

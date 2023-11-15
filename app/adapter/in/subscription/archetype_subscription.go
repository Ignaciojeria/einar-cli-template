package subscription

import (
	"context"
	"encoding/json"

	"archetype/app/shared/archetype/container"
	einar "archetype/app/shared/archetype/pubsub"
	"archetype/app/shared/archetype/pubsub/subscription"

	"archetype/app/shared/archetype/slog"

	"archetype/app/shared/constants"

	"time"

	"cloud.google.com/go/pubsub"
	"github.com/rs/zerolog/log"
)

var __archetype_subscription_stop bool = false

type __archetype_subscription_struct struct {
	subscriptionName string
}

func __archetype_subscription_constructor(
	r subscription.Receive,
	subscriptionName string) (__archetype_subscription_struct, error) {
	if __archetype_subscription_stop {
		return __archetype_subscription_struct{}, nil
	}
	s := __archetype_subscription_struct{
		subscriptionName: subscriptionName,
	}
	ctx := context.Background()
	if err := r(ctx, subscription.Middleware(subscriptionName, s.receive)); err != nil {
		log.
			Error().
			Err(err).
			Str(constants.SUBSCRIPTION_NAME, subscriptionName).
			Msg(constants.SUSBCRIPTION_SIGNAL_BROKEN)
		time.Sleep(10 * time.Second)
		go __archetype_subscription_constructor(r, subscriptionName)
		return s, err
	}
	return s, nil
}

func init() {
	const subscription_name = "INSERT YOUR SUBSCRIPTION NAME"
	container.InjectInboundAdapter(func() error {
		subscription_setup := einar.Client.Subscription(subscription_name)
		subscription_setup.ReceiveSettings.MaxOutstandingMessages = 5
		go __archetype_subscription_constructor(subscription_setup.Receive, subscription_name)
		return nil
	})
}

func (s __archetype_subscription_struct) receive(ctx context.Context, m *pubsub.Message) {
	s.processMessage(ctx, m)
}

func (s __archetype_subscription_struct) processMessage(ctx context.Context, m *pubsub.Message) (err error) {

	defer func() {
		if err != nil {
			slog.Logger.Error(
				"__archetype_subscription_error_when_pull",
				constants.SUBSCRIPTION_NAME, s.subscriptionName,
				constants.ERROR, err.Error(),
			)
			return
		}
		slog.Logger.Info(
			"__archetype_subscription_pull_succedded",
			constants.SUBSCRIPTION_NAME, s.subscriptionName,
			constants.ERROR, err.Error(),
		)
	}()

	var replace_by_your_model interface{}
	if json.Unmarshal(m.Data, &replace_by_your_model) != nil {
		m.Ack()
		return err
	}
	m.Ack()
	return nil
}

package subscription

import (
	"archetype/app/shared/archetype/slog"
	"archetype/app/shared/constants"
	"context"
	"time"

	"cloud.google.com/go/pubsub"
)

type Receive func(ctx context.Context, f func(context.Context, *pubsub.Message)) error

type Subscription struct {
	subscriptionName string
	processMessage   func(ctx context.Context, m *pubsub.Message) (err error)
	stop             bool
}

func New(
	subscriptionName string,
	processMessage func(ctx context.Context, m *pubsub.Message) (err error),
	recieveWithSettings Receive) (Subscription, error) {
	s := Subscription{
		subscriptionName: subscriptionName,
	}
	if s.stop {
		return Subscription{}, nil
	}
	ctx := context.Background()
	if err := recieveWithSettings(ctx, Middleware(subscriptionName, s.receive)); err != nil {
		slog.Logger.Error(
			constants.SUSBCRIPTION_SIGNAL_BROKEN,
			constants.SUBSCRIPTION_NAME, s.subscriptionName,
			constants.ERROR, err.Error(),
		)
		time.Sleep(10 * time.Second)
		go New(subscriptionName, processMessage, recieveWithSettings)
		return s, err
	}
	return s, nil
}

func (s Subscription) receive(ctx context.Context, m *pubsub.Message) {
	s.processMessage(ctx, m)
}

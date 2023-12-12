package subscription

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"archetype/app/shared/archetype/container"
	"archetype/app/shared/archetype/exception"
	archetype "archetype/app/shared/archetype/pubsub"
	"archetype/app/shared/archetype/pubsub/subscription"
	"archetype/app/shared/config"
	"archetype/app/shared/constants"
	"archetype/app/shared/ddlog"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"time"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
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
	container.InjectInBoundAdapter(func() error {
		subscription_name := "INSERT YOUR SUBSCRIPTION NAME"
		subscription_setup := archetype.Client.Subscription(subscription_name)
		subscription_setup.ReceiveSettings.NumGoroutines = 1
		subscription_setup.ReceiveSettings.MaxOutstandingMessages = 1
		go __archetype_subscription_constructor(subscription_setup.Receive, subscription_name)
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

func (s __archetype_subscription_struct) receive(ctx context.Context, m *pubsub.Message) {
	s.processMessage(ctx, m)
}

func (s __archetype_subscription_struct) processMessage(ctx context.Context, m *pubsub.Message) (err error) {
	span, ctx := tracer.StartSpanFromContext(ctx, s.subscriptionName, tracer.AnalyticsRate(1))
	var replace_by_your_model interface{}
	defer func() {
		if err != nil {
			span.SetTag("error.msg", err.Error())
			ddlog.
				Err(span.Context()).
				Interface("replace_by_your_model", replace_by_your_model).
				Err(err).
				Msg("INSERT_YOUR_CUSTOM_LOG_METRIC_ERROR")
			span.Finish(tracer.WithError(err))
			if errors.Is(err, exception.InternalServerError{}) {
				m.Nack()
				return
			}
			m.Ack()
			return
		}
		ddlog.
			Info(span.Context()).
			Interface("replace_by_your_model", replace_by_your_model).
			Msg("INSERT_YOUR_CUSTOM_LOG_METRIC_SUCCEDDED")
		m.Ack()
		span.Finish()
	}()

	if strings.ToLower(m.Attributes["country"]) != strings.ToLower(config.COUNTRY.Get()) {
		m.Ack()
		return exception.InvalidCountry{}
	}

	if err = json.Unmarshal(m.Data, &replace_by_your_model); err != nil {
		return err
	}

	return nil
}

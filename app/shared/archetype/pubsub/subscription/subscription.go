package subscription

import (
	"archetype/app/shared/archetype/echo_server"
	einar "archetype/app/shared/archetype/pubsub"
	"archetype/app/shared/archetype/slog"
	"archetype/app/shared/constants"
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/labstack/echo/v4"
)

const subscription_signal_broken = "subscription_signal_broken"
const subscription_name = "subscription_name"

type Receive func(ctx context.Context, f func(context.Context, *pubsub.Message)) error

type Subscription struct {
	subscriptionName    string
	processMessage      func(ctx context.Context, sub string, m *pubsub.Message) (statusCode int, err error)
	recieveWithSettings Receive
	stop                bool
}

func New(
	subscriptionName string,
	processMessage func(ctx context.Context, sub string, m *pubsub.Message) (statusCode int, err error),
	recieveWithSettings Receive) Subscription {
	return Subscription{subscriptionName: subscriptionName, processMessage: processMessage, recieveWithSettings: recieveWithSettings}
}

func (s Subscription) Start() (Subscription, error) {

	if s.stop {
		return Subscription{}, nil
	}

	if einar.Client().Project() == "" {
		s.recieveWithSettings = func(ctx context.Context, f func(context.Context, *pubsub.Message)) error {
			s.stop = true
			return errors.New("subscription cannot start")
		}
	}

	ctx := context.Background()

	if err := s.recieveWithSettings(ctx, s.receive); err != nil {
		slog.Logger().Error(
			subscription_signal_broken,
			subscription_name, s.subscriptionName,
			constants.Error, err.Error(),
		)
		time.Sleep(10 * time.Second)
		go s.Start()
		return s, err
	}
	return s, nil
}

func (s Subscription) WithPushHandler(path string) Subscription {
	echo_server.Echo().POST(path, s.pushHandler)
	return s
}

func (s Subscription) receive(ctx context.Context, m *pubsub.Message) {
	s.processMessage(ctx, s.subscriptionName, m)
}

func (s Subscription) pushHandler(c echo.Context) error {
	googleChannel := c.Request().Header.Get("X-Goog-Channel-ID")

	if googleChannel != "" {
		var msg pubsub.Message
		if err := c.Bind(&msg); err != nil {
			return c.String(http.StatusNoContent, "error binding Pub/Sub message")
		}
		statusCode, err := s.processMessage(c.Request().Context(), s.subscriptionName, &msg)
		if statusCode >= 500 && statusCode <= 599 {
			return c.String(statusCode, "")
		}
		if err != nil {
			return c.String(http.StatusNoContent, "")
		}
		return c.String(http.StatusOK, "")
	}

	if googleChannel == "" {
		var msg pubsub.Message
		msg.Attributes = map[string]string{
			constants.EventType:  c.Request().Header.Get(constants.EventType),
			constants.EntityType: c.Request().Header.Get(constants.EntityType),
		}
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.String(http.StatusBadRequest, "error reading request body")
		}
		msg.Data = body
		if statusCode, err := s.processMessage(c.Request().Context(), s.subscriptionName, &msg); err != nil {
			return c.String(statusCode, err.Error())
		}
		return c.String(http.StatusOK, "")
	}

	return c.String(http.StatusBadRequest, "unknown channel")
}

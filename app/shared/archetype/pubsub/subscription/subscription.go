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
	processMessage      func(ctx context.Context, sub string, m *pubsub.Message) (err error)
	recieveWithSettings Receive
	stop                bool
}

func New(
	subscriptionName string,
	processMessage func(ctx context.Context, sub string, m *pubsub.Message) (err error),
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

	if err := s.recieveWithSettings(ctx, Middleware(s.subscriptionName, s.receive)); err != nil {
		slog.Logger.Error(
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
	var msg pubsub.Message
	googleChannel := c.Request().Header.Get("X-Goog-Channel-ID")
	if googleChannel != "" {
		if err := c.Bind(&msg); err != nil {
			return c.String(http.StatusNoContent, "error binding Pub/Sub message")
		}
	}
	if googleChannel == "" {
		msg.Attributes = map[string]string{
			constants.EventType:  c.Request().Header.Get(constants.EventType),
			constants.EntityType: c.Request().Header.Get(constants.EntityType),
		}
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.String(http.StatusNoContent, "error reading request body")
		}
		msg.Data = body
	}
	if err := s.processMessage(c.Request().Context(), s.subscriptionName, &msg); err != nil {
		return c.String(http.StatusNoContent, "error processing message")
	}
	return nil
}

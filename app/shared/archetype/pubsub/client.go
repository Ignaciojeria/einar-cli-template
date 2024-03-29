package pubsub

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/archetype/slog"
	"archetype/app/shared/config"
	"archetype/app/shared/constants"
	"context"

	"cloud.google.com/go/pubsub"
)

var (
	client *pubsub.Client
)

func init() {
	config.Installations.EnablePubSub = true
	container.InjectInstallation(func() error {
		projectId := config.GOOGLE_PROJECT_ID.Get()
		c, err := pubsub.NewClient(context.Background(), projectId)
		if err != nil {
			slog.Logger().Error("error getting pubsub client", constants.Error, err.Error())
			return err
		}
		client = c
		return nil
	})
}

func Client() *pubsub.Client {
	if client == nil {
		return &pubsub.Client{}
	}
	return client
}

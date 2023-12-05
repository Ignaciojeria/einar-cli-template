package redis

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/archetype/slog"
	"archetype/app/shared/config"
	"archetype/app/shared/constants"
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func init() {
	config.Installations.EnableRedis = true
	container.InjectInstallation(func() error {
		addr := config.REDIS_ADDRESS.Get()
		db, err := strconv.Atoi(config.REDIS_DB.Get())
		if err != nil {
			fmt.Println(err)
			return err
		}
		Client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: config.REDIS_PASSWORD.Get(),
			DB:       db,
		})

		ping := Client.Ping(context.Background())
		if err := ping.Err(); err != nil {
			slog.Logger.Error("error on ping redis connection", constants.ERROR, err.Error())
			return err
		}
		return nil
	})
}

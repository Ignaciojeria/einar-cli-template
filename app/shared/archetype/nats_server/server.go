package nats_server

import (
	"archetype/app/shared/archetype/container"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

var Conn *nats.Conn

func init() {
	container.InjectInstallation(func() error {
		opts := &server.Options{}
		ns, err := server.NewServer(opts)
		if err != nil {
			return err
		}
		ns.Start()
		nc, err := nats.Connect(ns.ClientURL())
		if err != nil {
			return err
		}
		Conn = nc
		return err
	})
}

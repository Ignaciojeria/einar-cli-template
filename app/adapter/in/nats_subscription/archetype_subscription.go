package nats_subscription

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/archetype/nats_server"
	"fmt"

	"github.com/nats-io/nats.go"
)

func init() {
	container.InjectInboundAdapter(func() error {

		subjectName := "INSERT YOUR SUBJECT NAME HERE"

		processMessage := func(msg *nats.Msg) {
			//HANDLE MESSAGE PROCESSING HERE
			data := string(msg.Data)
			fmt.Println(data)
		}

		nats_server.Conn.Subscribe(subjectName, processMessage)
		return nil
	})
}

package nats_publisher

import (
	"archetype/app/shared/archetype/nats_server"
	"context"
	"encoding/json"
)

var ArchetypePublisher = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) (err error) {
	b, _ := json.Marshal(REPLACE_BY_YOUR_DOMAIN)
	nats_server.Conn.Publish("INSERT YOUR SUBJECT NAME HERE", b)
	return nil
}

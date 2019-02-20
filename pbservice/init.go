package pbservice

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"github.com/rs/zerolog/log"
)

var NatsConn *nats.Conn

func init() {
	NatsConn, _ = nats.Connect(nats.DefaultURL)

	go natsServer()
}

func natsServer() {

	for subkey, fun := range ServiceConfig {
		var fun = fun // go trap
		NatsConn.Subscribe(subkey, func(m *nats.Msg) {
			out, err := fun(m.Data)
			if err != nil {
				log.Info().Msg(fmt.Sprintf("%v", err.Error()))
			}
			NatsConn.Publish(m.Reply, out)
		})
	}
	println("Nats Server init done!")
	select {}
}

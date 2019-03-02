package natsqueue

import (
	"github.com/nats-io/go-nats"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("nats.client.url", nats.DefaultURL)
}

func NewClient() *nats.Conn {
	nc, err := nats.Connect(viper.GetString("nats.client.url"))
	if err != nil {
		log.Panic(err)
	}

	return nc
}

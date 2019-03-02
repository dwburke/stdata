package natsqueue

import (
	"os"

	"github.com/dwburke/atexit"
	stand "github.com/nats-io/nats-streaming-server/server"
	"github.com/nats-io/nats-streaming-server/stores"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("nats.server.enabled", false)
	viper.SetDefault("nats.server.port", 4222)
	viper.SetDefault("nats.server.filestoredir", ".nats-data-store")
}

type NatsServer struct {
	server *stand.StanServer
}

var NatsInstance *NatsServer

func Run() {
	if !viper.GetBool("nats.server.enabled") {
		return
	}

	NatsInstance = &NatsServer{}

	opts := stand.GetDefaultOptions()
	opts.StoreType = stores.TypeFile

	nopts := stand.NewNATSOptions()
	nopts.Port = viper.GetInt("nats.server.port")
	opts.FilestoreDir = os.ExpandEnv(viper.GetString("nats.server.filestoredir"))

	log.WithFields(log.Fields{
		"nats.server.filestoredir": opts.FilestoreDir,
		"nats.server.port":         nopts.Port,
		"StoreType":                opts.StoreType,
	}).Info("natsqueue: starting server")

	var err error
	NatsInstance.server, err = stand.RunServerWithOpts(opts, nopts)
	if err != nil {
		log.Panic(err)
	}

	atexit.Add(Shutdown)
}

func (ns *NatsServer) Shutdown() {
	if ns.server != nil {
		ns.server.Shutdown()
	}
}

func Shutdown() {
	log.Info("natsqueue: Shutdown()")
	if NatsInstance != nil {
		NatsInstance.Shutdown()
		NatsInstance = nil
	}
}

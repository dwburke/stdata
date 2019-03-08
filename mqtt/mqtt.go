package mqtt

import (
	"fmt"
	"log"
	//"net/url"
	//"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("client.mqtt.enabled", false)
	viper.SetDefault("client.mqtt.host", "localhost:1883")
	viper.SetDefault("client.mqtt.username", "")
	viper.SetDefault("client.mqtt.password", "")

	//viper.SetDefault("client.mqtt.url", "mqtt://localhost:<port>/<topic>");
	//mqtt://<user>:<pass>@<server>.cloudmqtt.com:<port>/<topic>
}

func Connect(clientId string) mqtt.Client {
	opts := createClientOptions(clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", viper.GetString("client.mqtt.host")))

	if viper.GetString("client.mqtt.username") != "" {
		opts.SetUsername(viper.GetString("client.mqtt.username"))
	}
	if viper.GetString("client.mqtt.password") != "" {
		opts.SetPassword(viper.GetString("client.mqtt.password"))
	}

	opts.SetClientID(clientId)
	return opts
}

func Listen(topic string, sub func(client mqtt.Client, msg mqtt.Message)) {
	client := Connect("sub")
	client.Subscribe(topic, 0, sub)
}

//func Listen(topic string) {
//client := Connect("sub")
//client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
//fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
//})
//}

// go Listen(topic)

// go Listen(topic)

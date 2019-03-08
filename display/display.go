package display

import (
	"fmt"
	"time"
)

func init() {
	//viper.SetDefault("client.mqtt.url", "mqtt://localhost:<port>/<topic>");
}

func Run() {
	for {
		fmt.Println(time.Now().Format("20060102150405"))
		fmt.Println(time.Now().Format("20060102150405"))

		select {
		case <-time.After(1 * time.Second):
			//case <-stop:
			//return
		}
	}

}

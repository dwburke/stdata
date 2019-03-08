package display

import (
	"fmt"
	//"log"
	//"net/url"
	//"os"
	"time"
	//mqtt "github.com/eclipse/paho.mqtt.golang"
	//"github.com/spf13/viper"
)

func init() {
	//viper.SetDefault("client.mqtt.url", "mqtt://localhost:<port>/<topic>");
	//mqtt://<user>:<pass>@<server>.cloudmqtt.com:<port>/<topic>
}

func Run() {
	//t := time.Now()
	//fmt.Println(t.Format("20060102150405"))

	for {
		fmt.Println(time.Now().Format("20060102150405"))
		select {
		case <-time.After(1 * time.Second):
		//case <-stop:
			//return
		}
	}

}

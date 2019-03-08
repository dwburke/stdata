package display

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/spf13/viper"

	"github.com/dwburke/stdata/db"
)

var stop chan bool
var refresh chan bool

func init() {
	//viper.SetDefault("client.mqtt.url", "mqtt://localhost:<port>/<topic>");
}

func Run() {

	stop = make(chan bool)
	refresh = make(chan bool)

	for {

		t := template.New("display.tmpl").Funcs(template.FuncMap{
			"viperGetString": func(key string) string {
				return viper.GetString(key)
			},
			"DoorLock": func(key string) string {
				value := db.LevelDBGet("topic:" + key)
				if value == "locked" {
					return "L"
				} else if value == "unlocked" {
					return "U"
				} else if value == "" {
					return "?"
				}
				return value
			},
			"DoorState": func(key string) string {
				value := db.LevelDBGet("topic:" + key)
				if value == "open" {
					return "O"
				} else if value == "closed" {
					return "C"
				} else if value == "" {
					return "?"
				}
				return value
			},
		})

		t2, err := t.ParseFiles("display.tmpl") //setp 1
		if err != nil {
			panic(err)
		}

		args := map[string]interface{}{
			"timestamp": time.Now().Format("Mon Jan 2 15:04:05 MST 2006"),
		}

		if err := t2.Execute(os.Stdout, args); err != nil {
			panic(err)
		}

		select {

		case <-stop:
			fmt.Println("Stop requested")
			return

		case <-refresh:
			// forced refresh

		case <-time.After(10 * time.Second):
			//timeout refresh
		}
	}

}

func Stop() {
	stop <- true
}

func Refresh() {
	refresh <- true
}

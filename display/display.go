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
				if value == "unlocked" {
					return "\033[0;31mU\033[0m"
				} else if value == "locked" {
					return "\033[0;32mL\033[0m"
				} else if value == "" {
					return "?"
				}
				return value
			},
			"DoorState": func(key string) string {
				value := db.LevelDBGet("topic:" + key)
				if value == "open" {
					return "\033[0;31mO\033[0m"
				} else if value == "closed" {
					return "\033[0;32mC\033[0m"
				} else if value == "" {
					return "?"
				}
				return value
			},
			"Light": func(key string) string {
				value := db.LevelDBGet("topic:" + key)
				if value == "on" {
					return "\033[0;107m*\033[0m"
				} else if value == "off" {
					return "-"
				} else if value == "" {
					return "?"
				}
				return value
			},
			"GetInt": func(key string) string {
				return db.LevelDBGet("topic:" + key)
			},
		})

		t2, err := t.ParseFiles("display.tmpl") //setp 1
		if err != nil {
			panic(err)
		}

		args := map[string]interface{}{
			//"timestamp": time.Now().Format("Mon Jan 2 15:04:05 MST 2006"),
		}

		f, err := os.Create("/tmp/display.out")
		if err != nil {
			panic(err)
		}

		if err := t2.Execute(f, args); err != nil {
			panic(err)
		}

		f.Close()

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

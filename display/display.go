package display

import (
	//"fmt"
	"os"
	"text/template"
	"time"

	"github.com/spf13/viper"

	"github.com/dwburke/stdata/db"
)

func init() {
	//viper.SetDefault("client.mqtt.url", "mqtt://localhost:<port>/<topic>");
}

func Run() {
	for {

		t := template.New("display.tmpl").Funcs(template.FuncMap{
			"viperGetString": func(key string) string {
				return viper.GetString(key)
			},
			"DoorLock": func(key string) string {
				value := db.LevelDBGet("topic:" + key)
				return value
			},
			"DoorState": func(key string) string {
				value := db.LevelDBGet("topic:" + key)
				return value
			},
		})

		t2, err := t.ParseFiles("display.tmpl") //setp 1
		if err != nil {
			panic(err)
		}

		//t2.Execute(w, "Hello World!") //step 2

		args := map[string]interface{}{
			"timestamp": time.Now().Format("Mon Jan 2 15:04:05 MST 2006"),
			//"timestamp": time.Now().Format("20060102150405"),
			//"Agent": agent,
		}

		if err := t2.Execute(os.Stdout, args); err != nil {
			panic(err)
		}

		//fmt.Println(time.Now().Format("20060102150405"))
		//fmt.Println("================\n")
		//fmt.Print(  "Front   : ")
		//fmt.Print(  "Garage  : ")
		//fmt.Print(  "Sun Room: ")
		//fmt.Print(  "Patio   : ")
		//fmt.Print(  "Office  : ")

		select {
		case <-time.After(1 * time.Second):
			//case <-stop:
			//return
		}
	}

}

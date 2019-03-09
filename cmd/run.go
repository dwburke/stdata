package cmd

import (
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/streadway/simpleuuid"

	"github.com/dwburke/stdata/db"
	"github.com/dwburke/stdata/display"
	stmqtt "github.com/dwburke/stdata/mqtt"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run",
	Long:  "run",
	Run: func(cmd *cobra.Command, args []string) {

		go stmqtt.Listen("smartthings/#", func(client mqtt.Client, msg mqtt.Message) {
			uuid, err := nextId()
			if err != nil {
				panic(err)
			}

			log.Printf("- [%s] [%s] %s\n", uuid, msg.Topic(), string(msg.Payload()))

			db.LevelDBSet("topic:"+msg.Topic(), string(msg.Payload()))
			db.LevelDBSet(uuid+":"+msg.Topic(), string(msg.Payload()))
			display.Refresh()
		})

		go display.Run()

		done := make(chan bool)
		<-done
	},
}

func nextId() (string, error) {
	uuid, err := simpleuuid.NewTime(time.Now())
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}

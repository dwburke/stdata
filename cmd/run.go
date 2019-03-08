package cmd

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

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
			log.Printf("- [%s] %s\n", msg.Topic(), string(msg.Payload()))
			db.LevelDBSet("topic:"+msg.Topic(), string(msg.Payload()))
		})

		go display.Run()

		done := make(chan bool)
		<-done
	},
}

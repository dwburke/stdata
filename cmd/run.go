package cmd

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

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

		//client mqtt.Client, msg mqtt.Message

		stmqtt.Listen("smartthings/#", func(client mqtt.Client, msg mqtt.Message) {
			log.Printf("- [%s] %s\n", msg.Topic(), string(msg.Payload()))
		})

		done := make(chan bool)
		<-done
	},
}

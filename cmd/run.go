package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "main entry point for running all of the enabled services",
	Long:  "main entry point for running all of the enabled services",
	Run: func(cmd *cobra.Command, args []string) {

		var started_something bool

		// go something.Run()
		// started_something = true

		if started_something {
			done := make(chan bool)
			<-done
		} else {
			log.Println("Nothing started... exiting")
		}
	},
}

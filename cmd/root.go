package cmd

import (
	"fmt"
	"os"

	"github.com/dwburke/cron"
	"github.com/dwburke/vipertools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/addictmud/mud/natsqueue"
)

var rootCmd = &cobra.Command{
	Use:   "mud",
	Short: "mud",
	Long:  `Main entrypoint`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var configList string

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(natsqueue.Run)
	cobra.OnInitialize(cron.Run)

	rootCmd.PersistentFlags().StringVar(&configList, "config-list", "etc/*.yml", "config files location")

	rootCmd.PersistentFlags().Bool("nats-server-enabled", false, "Override nats.server.enabled config value")
	viper.BindPFlag("nats.server.enabled", rootCmd.PersistentFlags().Lookup("nats-server-enabled"))
}

func initConfig() {
	if err := vipertools.ReadPattern(configList); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

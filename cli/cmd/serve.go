package cmd

import (
	"github.com/spf13/cobra"

	"github.com/protonhq/proton/delivery/http"
	"github.com/protonhq/proton/infra/config"
	"github.com/protonhq/proton/registry"
	log "github.com/sirupsen/logrus"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start proton server",
	Long:  `Start proton server`,
	Run: func(cmd *cobra.Command, args []string) {
		appConfig := config.NewConfiguration()
		_, err := registry.NewContainer(appConfig)

		if err != nil {
			log.Error("Error on initalizing application", err)
		}

		log.Info("Proton Server Version ", appConfig.Version)
		http.Init(appConfig.Server.Port)
	},
}

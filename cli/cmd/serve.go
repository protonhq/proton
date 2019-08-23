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
		appContainer, err := registry.NewContainer()
		if err != nil {
			log.Error("Error on initalizing application: ", err)
		}
		appConfig := appContainer.Resolve("config").(*config.Configuration)

		log.Info("Proton Server Version ", appConfig.Version)
		http.Init(appContainer)
	},
}

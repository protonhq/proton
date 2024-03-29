package cmd

import (
	"github.com/protonhq/proton/infra/config"
	"github.com/protonhq/proton/infra/db"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		appConfig := config.NewConfiguration()
		database, err := db.InitDB(appConfig.Database.ConnectionString())

		if err != nil {
			log.Error(err)
		}

		db.Migrate(database)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

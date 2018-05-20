package cmd

import (
	"log"

	"github.com/chuckleheads/hurtlocker/components/sessionsrv/data_store"
	"github.com/chuckleheads/hurtlocker/components/sessionsrv/data_store/migrations"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run the database migrations for sessionsrv",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Running Migrations...")
		dbConfig, err := DBConfigFromViper()
		if err != nil {
			panic(err.Error())
		}
		db := data_store.New(dbConfig)

		migrations.Migrate(db, dbConfig.Migrations)
		log.Printf("Migrations complete")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

package cmd

import (
	"log"

	"github.com/chuckleheads/hurtlocker/components/jobsrv/data_store"
	"github.com/chuckleheads/hurtlocker/components/jobsrv/data_store/migrations"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run the database migrations for jobsrv",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Running Migrations...")
		config, err := ConfigFromViper()
		if err != nil {
			panic(err.Error())
		}
		db := data_store.New(&config.Datastore)

		migrations.Migrate(db, config.Datastore.Migrations)
		log.Printf("Migrations complete")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

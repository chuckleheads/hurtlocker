package cmd

import (
	"log"
	"os"

	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/export"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// exportCmd represents the start command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Execute a Habitat package",
	Run: func(cobraCmd *cobra.Command, args []string) {
		exportCli := export.New()
		exportCli.Export()
	},
}

func init() {
	cobra.OnInitialize(validateExportConfig)
	rootCmd.AddCommand(exportCmd)
}

// validateConfig ensures any config values we depend on for running a job are present before starting
func validateExportConfig() {
	requiredConfig := []string{"ident",
		"channel",
	}
	for _, req := range requiredConfig {
		if !viper.IsSet(req) {
			log.Fatalln("Missing required config value for: ", req)
			os.Exit(1)
		}
	}
}

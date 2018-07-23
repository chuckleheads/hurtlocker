package cmd

import (
	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/export"
	"github.com/spf13/cobra"
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
	rootCmd.AddCommand(exportCmd)
}

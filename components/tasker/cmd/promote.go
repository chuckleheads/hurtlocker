package cmd

import (
	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/promote"
	"github.com/spf13/cobra"
)

// promoteCmd represents the start command
var promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "Promote a Habitat package",
	Run: func(cobraCmd *cobra.Command, args []string) {
		promoteCli := promote.New()
		promoteCli.Promote()
	},
}

func init() {
	rootCmd.AddCommand(promoteCmd)
}

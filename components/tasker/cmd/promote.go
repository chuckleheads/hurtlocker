package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/cmd"
	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/promote"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// buildCmd represents the start command
var promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "Promote a Habitat package",
	Run: func(cobraCmd *cobra.Command, args []string) {
		client := cmd.LogSrvClient()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		stream, err := client.ReceiveLogs(ctx)
		if err != nil {
			log.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
		}
		for k, v := range viper.GetStringMap("stages") {
			fmt.Printf("Key: %s, %s\n", k, v)
		}
		promoteCli := promote.New(stream)
		promoteCli.Promote()
	},
}

func init() {
	rootCmd.AddCommand(promoteCmd)
}

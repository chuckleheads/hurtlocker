package cmd

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/cmd"
// 	"github.com/spf13/cobra"
// )

// // exportCmd represents the start command
// var exportCmd = &cobra.Command{
// 	Use:   "export",
// 	Short: "Execute a Habitat package",
// 	Run: func(cobraCmd *cobra.Command, args []string) {
// 		client := cmd.LogSrvClient()
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()
// 		stream, err := client.ReceiveLogs(ctx)
// 		if err != nil {
// 			log.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
// 		}
// 		// exportCli := build.New(stream, dir, viper.GetString("repo_url"), viper.GetString("plan_path"))
// 		// exportCli.Start()
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(exportCmd)
// }

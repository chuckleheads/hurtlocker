package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/build"
	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// buildCmd represents the start command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Execute a Habitat build",
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
		dir := tempdir()
		// defer os.RemoveAll(dir)
		buildCli := build.New(stream, dir, viper.GetString("repo_url"), viper.GetString("plan_path"))
		buildCli.Start()
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func tempdir() string {
	dir, err := ioutil.TempDir("", "hab-root")
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

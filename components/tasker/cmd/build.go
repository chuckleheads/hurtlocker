package cmd

import (
	"io/ioutil"
	"log"

	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/build"
	"github.com/spf13/cobra"
)

// buildCmd represents the start command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Execute a Habitat build",
	Run: func(cobraCmd *cobra.Command, args []string) {
		dir := tempdir()
		buildCli := build.New(dir)
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

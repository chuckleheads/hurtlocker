package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/chuckleheads/hurtlocker/components/tasker/pkg/build"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	cobra.OnInitialize(validateBuildConfig)
	rootCmd.AddCommand(buildCmd)
}

func tempdir() string {
	dir, err := ioutil.TempDir("", "hab-root")
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

// validateConfig ensures any config values we depend on for running a job are present before starting
func validateBuildConfig() {
	requiredConfig := []string{"project.origin_name",
		"project.package_name",
		"project.name",
		"project.plan_path",
		"project.vcs_data",
		"bldr.private",
		"bldr.public",
	}
	for _, req := range requiredConfig {
		if !viper.IsSet(req) {
			log.Fatalln("Missing required config value for: ", req)
			os.Exit(1)
		}
	}
}

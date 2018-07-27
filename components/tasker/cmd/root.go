package cmd

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var cfgString string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tasker",
	Short: "A task runner that executes tasks and streams the output back to the API",
}

func init() {
	cobra.OnInitialize(initConfig, validateConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.PersistentFlags().StringVar(&cfgString, "config-string", "", "a base64 encoded config string (useful for passing config though containers)")
	rootCmd.PersistentFlags().Bool("enable-log-stream", false, "Enable log streaming via gRPC")
	viper.BindPFlag("enable-log-stream", rootCmd.Flags().Lookup("enable-log-stream"))
	viper.SetDefault("bldr-url", "https://bldr.habitat.sh")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			log.Fatal(err)
		}
	} else {
		viper.SetConfigType("json")
		cfg, err := base64.StdEncoding.DecodeString(cfgString)
		if err != nil {
			log.Fatalln(err)
		}
		err = viper.ReadConfig(bytes.NewBuffer(cfg))
		if err != nil {
			log.Fatalln(err)
		}
	}

	viper.AutomaticEnv() // read in environment variables that match
}

func validateConfig() {
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

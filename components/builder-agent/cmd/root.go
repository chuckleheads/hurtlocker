package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/chuckleheads/hurtlocker/components/builder-agent/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "builder-agent",
	Short: "An agent that watches the builder scheduler queue for work",
	Long: `Builder Agent watches the scheduler work queue and translates payloads
into runnable workloads for Builder Tasker
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /hab/svc/builder-agent/config/config.toml)")
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 5672)
	viper.SetDefault("username", "guest")
	viper.SetDefault("password", "guest")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home := filepath.Dir("/hab/svc/builder-agent/config")

		// Search config in home directory with name ".config" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// RabbitMQConfigFromViper fetches database config from viper
func RabbitMQConfigFromViper() (*config.RabbitMQConfig, error) {
	cfg := &config.RabbitMQConfig{}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(err.Error())
	}
	return cfg, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

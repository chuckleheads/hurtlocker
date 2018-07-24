package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/chuckleheads/hurtlocker/components/agent/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "agent",
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /hab/svc/agent/config/config.toml)")
	viper.SetDefault("driver", "docker")
	viper.SetDefault("rabbitmq.host", "localhost")
	viper.SetDefault("rabbitmq.port", 5672)
	viper.SetDefault("rabbitmq.username", "guest")
	viper.SetDefault("rabbitmq.password", "guest")
	viper.SetDefault("rabbitmq.exchange", "work")
	viper.SetDefault("rabbitmq.queue", "linux")
	// Listen for all topics by default
	viper.SetDefault("rabbitmq.topic", []string{"linux.build", "linux.export"})

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home := filepath.Dir("/hab/svc/agent/config")

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

// ConfigFromViper fetches database config from viper
func ConfigFromViper() (*config.Config, error) {
	cfg := &config.Config{}
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

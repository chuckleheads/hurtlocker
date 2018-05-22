package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/chuckleheads/hurtlocker/components/sessionsrv/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sessionsrv",
	Short: "Session Server deals with accounts and authentication",
	Long:  `Session Server is teh best.`,
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /hab/svc/sessionsrv/config/config.toml)")
	viper.SetDefault("datastore.host", "localhost")
	viper.SetDefault("datastore.port", 26257)
	viper.SetDefault("datastore.database", "sessionsrv")
	viper.SetDefault("datastore.username", "root")
	viper.SetDefault("datastore.password", "")
	viper.SetDefault("datastore.ssl-mode", "disable")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home := filepath.Dir("/hab/svc/sessionsrv/config")

		// Search config in home directory with name "config" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Printf("Err: %v", err)
	}

}

// ConfigFromViper fetches database config from viper
func ConfigFromViper() (*config.Config, error) {
	cfg := &config.Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(err.Error())
	}
	log.Printf("CONFIG: %v", cfg)
	return cfg, nil
}

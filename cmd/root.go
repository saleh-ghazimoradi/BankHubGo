package cmd

import (
	"fmt"
	"github.com/saleh-ghazimoradi/BankHubGo/config"
	"github.com/saleh-ghazimoradi/BankHubGo/logger"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "BankHubGo",
	Short: "A simple bank management system",
	Long: `BankHubGo is a streamlined bank management service built with Go, 
designed to introduce core backend development topics and modern tools. 
This project covers the fundamentals required to build scalable backend services.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := os.Setenv("TZ", time.UTC.String())
	if err != nil {
		panic(err)
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file path (default is .)")
	cobra.OnInitialize(
		initConfig,
		initLogger, // Logger initialization comes after config
	)
}

func initConfig() {
	err := config.LoadingConfig(cfgFile)
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	fmt.Printf("Configuration loaded: %+v\n", config.Appconfig) // Verify config
}

func initLogger() {
	logger.LoadLogger(config.Appconfig.LogLevel)
}

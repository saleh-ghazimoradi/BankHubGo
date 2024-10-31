/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
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
	cfgFile   string
	appConfig config.Config // Global variable to store config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "BankHubGo",
	Short: "A simple bank management system",
	Long: `BankHubGo is a streamlined bank management service built with Go, 
designed to introduce core backend development topics and modern tools. 
This project covers the fundamentals required to build scalable backend services.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is .)")
	cobra.OnInitialize(
		initConfig,
		initLogger, // logger should come after config
	)
}

func initConfig() {
	var err error
	appConfig, err = config.LoadingConfig(cfgFile)
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	fmt.Printf("Configuration loaded: %+v\n", appConfig) // Print config for verification
}

func initLogger() {
	logger.LoadLogger(appConfig.LogLevel)
}

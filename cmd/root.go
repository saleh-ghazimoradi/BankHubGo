/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/saleh-ghazimoradi/BankHubGo/config"
	"github.com/saleh-ghazimoradi/BankHubGo/logger"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "BankHubGo",
	Short: "A simple bank management system",
	Long: `BankHubGo is a streamlined bank management service built with Go, 
designed to introduce core backend development topics and modern tools. 
This project covers the fundamentals required to build scalable backend services.
`,
}

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
	cobra.OnInitialize(
		initConfig,
		initLogger, // logger should come after config
	)
}

func initConfig() {
	err := config.LoadingConfig(".") // Provide the path to app.env
	if err != nil {
		logger.Logger.Fatal("failed to load configuration: ", err)
	}
}

func initLogger() {
	logger.LoadLogger(config.Appconfig.LogLevel)
}

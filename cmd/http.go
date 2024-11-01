/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/saleh-ghazimoradi/BankHubGo/config"
	"github.com/saleh-ghazimoradi/BankHubGo/internal/gateway"
	"github.com/saleh-ghazimoradi/BankHubGo/internal/repository"
	"github.com/saleh-ghazimoradi/BankHubGo/internal/service"
	"github.com/saleh-ghazimoradi/BankHubGo/logger"
	utils "github.com/saleh-ghazimoradi/BankHubGo/utils/connections"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "launching the http rest listen server",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Logger.Infow("server has started", "addr", config.Appconfig.ServerAddress, "env", config.Appconfig.Env)

		cfg := utils.PostConfig{
			Host:         config.Appconfig.DBHost,
			Port:         config.Appconfig.DBPort,
			User:         config.Appconfig.DBUser,
			Password:     config.Appconfig.DBPassword,
			Database:     config.Appconfig.DBName,
			SSLMode:      config.Appconfig.DBSSLMode,
			MaxIdleTime:  config.Appconfig.MaxIdleTime,
			MaxIdleConns: config.Appconfig.MaxIdleConns,
			MaxOpenConns: config.Appconfig.MaxOpenConns,
			Timeout:      config.Appconfig.Timeout,
		}

		db, err := utils.PostConnection(cfg)
		if err != nil {
			logger.Logger.Fatal(err)
		}
		defer db.Close()

		logger.Logger.Infow("Postgresql connection pool established")

		accountDB := repository.NewAccountRepository(db)
		accountService := service.NewAccountService(accountDB)
		accountHandler := gateway.NewAccountHandler(accountService)

		routeHandlers := gateway.Handlers{
			GetAccount:    accountHandler.GetAccount,
			GetAccounts:   accountHandler.GetAccounts,
			CreateAccount: accountHandler.CreateAccount,
			UpdateAccount: accountHandler.UpdateAccount,
			DeleteAccount: accountHandler.DeleteAccount,
		}

		if err := gateway.Server(gateway.Routes(routeHandlers)); err != nil {
			logger.Logger.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

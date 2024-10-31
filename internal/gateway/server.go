package gateway

import (
	"context"
	"errors"
	"github.com/saleh-ghazimoradi/BankHubGo/config"
	"github.com/saleh-ghazimoradi/BankHubGo/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Server(mux http.Handler) error {
	srv := http.Server{
		Addr:         config.Appconfig.ServerAddress,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		logger.Logger.Infow("signal caught", "signal", s.String())

		shutdown <- srv.Shutdown(ctx)
	}()

	err := srv.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	err = <-shutdown
	if err != nil {
		return err
	}

	logger.Logger.Infow("server has stopped", "addr", config.Appconfig.ServerAddress, "env", config.Appconfig.Env)

	return nil
}

package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aintsashqa/go-project-template/internal/config"
	httpapi "github.com/aintsashqa/go-project-template/internal/delivery/http-api"
	"github.com/aintsashqa/go-project-template/pkg/postgres"
	slogext "github.com/aintsashqa/go-project-template/pkg/slog-ext"
)

func main() {
	conf := config.Get()
	logger := slogext.New(conf.Logger)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	pool, err := postgres.NewPool(ctx, conf.Postgres)
	if err != nil {
		logger.ErrorContext(ctx, "failed to create postgres pool", slogext.Error(err))
		os.Exit(1)
	}

	defer pool.Close()

	srv := http.Server{
		Addr:           fmt.Sprintf(":%d", conf.HTTP.Port),
		Handler:        httpapi.NewHandler(),
		ReadTimeout:    conf.HTTP.ReadTimeout,
		WriteTimeout:   conf.HTTP.WriteTimeout,
		IdleTimeout:    conf.HTTP.IdleTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		logger.InfoContext(ctx, "starting http server", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorContext(ctx, "failed listen and serve", slogext.Error(err))
		}
	}()

	<-ctx.Done()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second*15)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.ErrorContext(shutdownCtx, "failed to shutdown http server", slogext.Error(err))
		os.Exit(1)
	}

	logger.Info("successfully shutdown")
}

// Package app handles starting and monitoring the server for graceful shutdown.
package app

import (
	"context"
	"fmt"

	"github.com/op/go-logging"

	"github.com/Egor123qwe/logs-storage/internal/server"
	"github.com/Egor123qwe/logs-storage/internal/service"
	"github.com/Egor123qwe/logs-storage/internal/storage"
)

var log = logging.MustGetLogger("app")

type App struct {
	srv service.Service
}

func New() (*App, error) {
	storage, err := storage.New()
	if err != nil {
		return nil, err
	}

	app := &App{
		srv: service.New(storage),
	}

	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	srv, err := server.New(a.srv)
	if err != nil {
		return err
	}

	if err := srv.Serve(ctx); err != nil {
		return fmt.Errorf("server stopped with error: %w\n", err)
	}

	log.Infof("server stopped")

	return nil
}

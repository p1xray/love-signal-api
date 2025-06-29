package app

import (
	"context"
	"errors"
	"log/slog"
	grpcapp "love-signal-api/internal/app/grpc"
	httpapp "love-signal-api/internal/app/http"
	"love-signal-api/internal/config"
	"love-signal-api/internal/lib/logger/sl"
	"net/http"
	"time"
)

// App is an application.
type App struct {
	log     *slog.Logger
	httpApp *httpapp.App
}

// New creates new instance of application.
func New(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	grpcApp := grpcapp.New(log, cfg)
	grpcClient, err := grpcApp.CreateGRPCClient()
	if err != nil {
		panic(err)
	}

	httpApp := httpapp.New(log, cfg.Server.Port, grpcClient)

	return &App{
		log:     log,
		httpApp: httpApp,
	}
}

// MustRun runs the application and panics if an error occurs.
func (a *App) MustRun() {
	if err := a.httpApp.Run(); !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

// GracefulStop stops the application.
func (a *App) GracefulStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.httpApp.Stop(ctx); err != nil {
		a.log.Error("HTTP app stop error", sl.Err(err))
	}
}

package httpapp

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
)

// App is an HTTP server application.
type App struct {
	log        *slog.Logger
	httpServer *http.Server
}

// New creates new instance of HTTP server application.
func New(log *slog.Logger, port int) *App {
	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		// TODO: Handler: handlers.Init(),
	}

	return &App{
		log:        log,
		httpServer: httpServer,
	}
}

// Run starts the server.
func (a *App) Run() error {
	const op = "httpapp.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.String("addr", a.httpServer.Addr),
	)

	log.Info("running HTTP server")

	if err := a.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Stop stops the server.
func (a *App) Stop(ctx context.Context) error {
	const op = "httpapp.Stop"

	log := a.log.With(
		slog.String("op", op),
		slog.String("addr", a.httpServer.Addr),
	)

	log.Info("shutdowning HTTP server")

	if err := a.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("HTTP server is shutdown")

	return nil
}

package grpcapp

import (
	"fmt"
	lsuserspb "github.com/p1xray/love-signal-protos/gen/go/users"
	urlshortenerpb "github.com/p1xray/pxr-url-shortener/pkg/grpc/gen/go/urlshortener"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	grpcclient "love-signal-api/internal/client/grpc"
	"love-signal-api/internal/config"
	"love-signal-api/internal/lib/logger/sl"
)

// App is gRPC client application.
type App struct {
	log    *slog.Logger
	config *config.Config
}

// New creates new instance of the gRPC client application.
func New(log *slog.Logger, cfg *config.Config) *App {
	return &App{
		log:    log,
		config: cfg,
	}
}

// CreateGRPCClient creates new gRPC clients.
func (a *App) CreateGRPCClient() (*grpcclient.GRPCClient, error) {
	users, err := a.createUsersClient()
	if err != nil {
		a.log.Error("error creating users grpc client", sl.Err(err))
		return nil, fmt.Errorf("error creating users grpc client: %w", err)
	}

	urlShortener, err := a.createUrlShortenerClient()
	if err != nil {
		a.log.Error("error creating url shortener grpc client", sl.Err(err))
		return nil, fmt.Errorf("error creating url shortener grpc client: %w", err)
	}

	return grpcclient.New(users, urlShortener), nil
}

func (a *App) createUsersClient() (lsuserspb.UsersClient, error) {
	const op = "grpcapp.createUsersClient"

	con, err := grpc.NewClient(
		a.config.GRPCClients.Users.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	usersClient := lsuserspb.NewUsersClient(con)
	return usersClient, nil
}

func (a *App) createUrlShortenerClient() (urlshortenerpb.UrlShortenerClient, error) {
	const op = "grpcapp.createUrlShortenerClient"

	con, err := grpc.NewClient(
		a.config.GRPCClients.UrlShortener.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	urlShortenerClient := urlshortenerpb.NewUrlShortenerClient(con)
	return urlShortenerClient, nil
}

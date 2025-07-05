package v1

import (
	"github.com/gin-gonic/gin"
	grpcclient "love-signal-api/internal/client/grpc"
	"love-signal-api/internal/config"
	"love-signal-api/internal/controller/http/v1/ping"
	"love-signal-api/internal/controller/http/v1/users"
)

// Handler is request handler for API v1.
type Handler struct {
	config     *config.Config
	grpcClient *grpcclient.GRPCClient
}

// New creates new instance of the API v1 request handler.
func New(cfg *config.Config, grpcClient *grpcclient.GRPCClient) *Handler {
	return &Handler{
		config:     cfg,
		grpcClient: grpcClient,
	}
}

// Init initializes the API v1 request handler.
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		ping.InitRoutes(v1)
		users.InitRoutes(v1, h.config.Server, h.grpcClient.Users, h.grpcClient.UrlShortener, h.grpcClient.QRCode)
	}
}

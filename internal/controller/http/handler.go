package http

import (
	"github.com/gin-gonic/gin"
	grpcclient "love-signal-api/internal/client/grpc"
	v1 "love-signal-api/internal/controller/http/v1"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "love-signal-api/docs"
)

// Handler is handler for http server requests.
type Handler struct {
	grpcClient *grpcclient.GRPCClient
}

// New creates a new http server request handler.
func New(grpcClient *grpcclient.GRPCClient) *Handler {
	return &Handler{grpcClient: grpcClient}
}

// Init initializes the http server request handler.
func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	h.initAPI(router)
	initSwagger(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	v1Handler := v1.New(h.grpcClient)
	api := router.Group("/api")
	{
		v1Handler.Init(api)
	}
}

func initSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

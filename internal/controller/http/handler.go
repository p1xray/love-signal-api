package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	grpcclient "love-signal-api/internal/client/grpc"
	"love-signal-api/internal/config"
	v1 "love-signal-api/internal/controller/http/v1"
	"love-signal-api/pkg/kafka"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "love-signal-api/docs"
)

// Handler is handler for http server requests.
type Handler struct {
	config        *config.Config
	grpcClient    *grpcclient.GRPCClient
	kafkaSendData chan<- kafka.Message
}

// New creates a new http server request handler.
func New(cfg *config.Config, grpcClient *grpcclient.GRPCClient, kafkaSendData chan<- kafka.Message) *Handler {
	return &Handler{
		config:        cfg,
		grpcClient:    grpcClient,
		kafkaSendData: kafkaSendData,
	}
}

// Init initializes the http server request handler.
func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AllowHeaders = []string{"Content-Type", "Authorization", "X-Fingerprint"}
	router.Use(cors.New(corsCfg))

	h.initAPI(router)
	initSwagger(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	v1Handler := v1.New(h.config, h.grpcClient, h.kafkaSendData)
	api := router.Group("/api")
	{
		v1Handler.Init(api)
	}
}

func initSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

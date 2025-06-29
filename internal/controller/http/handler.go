package http

import (
	"github.com/gin-gonic/gin"
	v1 "love-signal-api/internal/controller/http/v1"
)

// Handler is handler for http server requests.
type Handler struct {
}

// New creates a new http server request handler.
func New() *Handler {
	return &Handler{}
}

// Init initializes the http server request handler.
func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	v1Handler := v1.New()
	api := router.Group("/api")
	{
		v1Handler.Init(api)
	}
}

package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Routes provides routes for checking connectivity.
type Routes struct {
}

// InitRoutes initializes routes to checking connectivity.
func InitRoutes(api *gin.RouterGroup) {
	r := &Routes{}

	ping := api.Group("/ping")
	{
		ping.GET("", r.ping)
	}
}

func (r *Routes) ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

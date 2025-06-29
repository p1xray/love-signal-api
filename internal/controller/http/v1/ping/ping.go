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

// Checking the connection
//
//	@Summary		Checking the connection
//	@Description	Checking the connection
//	@Tags			Ping
//	@Id 			ping
//	@Accept			json
//	@Produce		plain
//	@Success		200	{string}	string	"pong"
//	@Failure		400	{string}	string	"ok"
//	@Failure		404	{string}	string	"ok"
//	@Failure		500	{string}	string	"ok"
//
// @Router			/api/v1/ping [get]
func (r *Routes) ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

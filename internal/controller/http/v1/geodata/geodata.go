package geodata

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"love-signal-api/internal/server"
	"love-signal-api/pkg/kafka"
)

// Routes provides routes for geodata.
type Routes struct {
	kafkaSendData chan<- kafka.Message
}

// InitRoutes initializes the routes for geodata.
func InitRoutes(
	api *gin.RouterGroup,
	kafkaSendData chan<- kafka.Message,
) {
	r := &Routes{
		kafkaSendData: kafkaSendData,
	}

	geodata := api.Group("/geodata")
	{
		geodata.POST("", r.processGeodata)
	}
}

// Process user's geodata.
//
//	@Summary		Process user's geodata
//	@Description	Process user's geodata
//	@Tags			Geodata
//	@Id 			processGeodata
//	@Produce		json
//	@Param			input body ProcessGeodataInput true "Input parameters for process user's geodata."
//	@Success		200	{object}  server.dataResponse[bool]
//	@Failure		500	{object}  server.dataResponse[bool]
//	@Router			/api/v1/geodata [post]
func (r *Routes) processGeodata(c *gin.Context) {
	inp, err := server.GetInputFromBody[ProcessGeodataInput](c)
	if err != nil {
		server.ErrorResponse[bool](c, err.Error())
		return
	}

	userGeodataAsBytes, err := json.Marshal(inp)
	if err != nil {
		server.ErrorResponse[bool](c, err.Error())
	}

	kafkaMessageToSend := kafka.Message{
		Topic: "user-geo-data",
		Data:  userGeodataAsBytes,
	}

	r.kafkaSendData <- kafkaMessageToSend

	success := true
	server.SuccessResponse(c, &success)
}

package users

import (
	"github.com/gin-gonic/gin"
	lsuserspb "github.com/p1xray/love-signal-protos/gen/go/users"
	"love-signal-api/internal/server"
)

// Routes provides routes for users.
type Routes struct {
	grpcUsersClient lsuserspb.UsersClient
}

// InitRoutes initializes the routes for users.
func InitRoutes(api *gin.RouterGroup, grpcUsersClient lsuserspb.UsersClient) {
	r := &Routes{grpcUsersClient: grpcUsersClient}

	profile := api.Group("/users")
	{
		profile.GET("/user-info", r.userInfo)
	}
}

// Current user info.
//
//	@Summary		Current user info
//	@Description	Current user info
//	@Tags			Users
//	@Id 			userInfo
//	@Produce		json
//	@Security 		ApiKeyAuth
//	@Success		200	{object}  server.dataResponse[UserInfoOutput]
//	@Failure		500	{object}  server.dataResponse[UserInfoOutput]
//	@Router			/api/v1/users/user-info [get]
func (r *Routes) userInfo(c *gin.Context) {
	userExternalID := int64(1) // TODO: get it from token

	grpcUserDataRequest := &lsuserspb.GetUserDataByExternalIdRequest{UserExternalId: userExternalID}
	grpcUserDataResponse, err := r.grpcUsersClient.GetUserDataByExternalId(c.Request.Context(), grpcUserDataRequest)
	if err != nil {
		// TODO: check error from gRPC server and return correct error

		server.ErrorResponse[UserInfoOutput](c, err.Error())
		return
	}

	userInfo := UserInfoOutput{
		ID:       grpcUserDataResponse.GetId(),
		FullName: grpcUserDataResponse.GetFullName(),
	}

	server.SuccessResponse(c, &userInfo)
}

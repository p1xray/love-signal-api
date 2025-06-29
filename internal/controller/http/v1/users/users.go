package users

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	lsuserspb "github.com/p1xray/love-signal-protos/gen/go/users"
	"love-signal-api/internal/server"
	"time"
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
		profile.GET("/profile", r.userProfileCard)
		profile.GET("/follow-link", r.followLinkCard)
		profile.GET("/follow/:id", r.followCard)
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

// User profile card.
//
//	@Summary		User profile card
//	@Description	User profile card
//	@Tags			Users
//	@Id 			userProfileCard
//	@Produce		json
//	@Security 		ApiKeyAuth
//	@Success		200	{object}  server.dataResponse[UserProfileCardOutput]
//	@Failure		500	{object}  server.dataResponse[UserProfileCardOutput]
//	@Router			/api/v1/users/profile [get]
func (r *Routes) userProfileCard(c *gin.Context) {
	userExternalID := int64(1) // TODO: get it from token

	grpcUserDataRequest := &lsuserspb.GetUserDataByExternalIdRequest{UserExternalId: userExternalID}
	grpcUserDataResponse, err := r.grpcUsersClient.GetUserDataByExternalId(c.Request.Context(), grpcUserDataRequest)
	if err != nil {
		// TODO: check error from gRPC server and return correct error

		server.ErrorResponse[UserProfileCardOutput](c, err.Error())
		return
	}

	var dateOfBirth *time.Time
	if grpcUserDataResponse.GetDateOfBirth() != nil {
		dateOfBirthValue := grpcUserDataResponse.GetDateOfBirth().AsTime()
		dateOfBirth = &dateOfBirthValue
	}

	var gender *GenderEnum
	if grpcUserDataResponse.GetGender() != lsuserspb.Gender_GENDER_UNSPECIFIED {
		genderValue := GenderEnum(grpcUserDataResponse.GetGender())
		gender = &genderValue
	}

	var avatarFileKey *string
	if grpcUserDataResponse.GetAvatarFileKey() != nil {
		avatarFileKeyValue := grpcUserDataResponse.GetAvatarFileKey().GetValue()
		avatarFileKey = &avatarFileKeyValue
	}

	userProfileCard := UserProfileCardOutput{
		ID:            grpcUserDataResponse.GetId(),
		FullName:      grpcUserDataResponse.GetFullName(),
		Gender:        gender,
		DateOfBirth:   dateOfBirth,
		AvatarFileKey: avatarFileKey,
	}

	server.SuccessResponse(c, &userProfileCard)
}

// User follow link card.
//
//	@Summary		User follow link card
//	@Description	User follow link card
//	@Tags			Users
//	@Id 			followLinkCard
//	@Produce		json
//	@Security 		ApiKeyAuth
//	@Success		200	{object}  server.dataResponse[UserFollowLinkCardOutput]
//	@Failure		500	{object}  server.dataResponse[UserFollowLinkCardOutput]
//	@Router			/api/v1/users/follow-link [get]
func (r *Routes) followLinkCard(c *gin.Context) {
	userExternalID := int64(1) // TODO: get it from token

	grpcUserDataRequest := &lsuserspb.GetUserDataByExternalIdRequest{UserExternalId: userExternalID}
	grpcUserDataResponse, err := r.grpcUsersClient.GetUserDataByExternalId(c.Request.Context(), grpcUserDataRequest)
	if err != nil {
		// TODO: check error from gRPC server and return correct error

		server.ErrorResponse[UserFollowLinkCardOutput](c, err.Error())
		return
	}

	host := server.GetHost(c)
	userID := grpcUserDataResponse.GetId()
	followLink := fmt.Sprintf("%s/api/v1/users/follow/%d", host, userID)

	// TODO: shorten link
	shortFollowLink := followLink

	// TODO: generate QR code from short follow link
	QRCode := base64.StdEncoding.EncodeToString([]byte("qr-code"))

	userFollowLinkCard := UserFollowLinkCardOutput{
		ShortLink: shortFollowLink,
		QRCode:    QRCode,
	}

	server.SuccessResponse(c, &userFollowLinkCard)
}

// User follow card.
//
//	@Summary		User follow card
//	@Description	User follow card
//	@Tags			Users
//	@Id 			followCard
//	@Produce		json
//	@Security 		ApiKeyAuth
//	@Param			id	path  int  true  "User ID"
//	@Success		200	{object}  server.dataResponse[UserFollowCardOutput]
//	@Failure		500	{object}  server.dataResponse[UserFollowCardOutput]
//	@Router			/api/v1/users/follow/{id} [get]
func (r *Routes) followCard(c *gin.Context) {
	userID, err := server.GetIdFromRoute(c)
	if err != nil {
		server.ErrorResponse[UserFollowCardOutput](c, err.Error())
		return
	}

	grpcUserDataRequest := &lsuserspb.GetUserDataRequest{UserId: userID}
	grpcUserDataResponse, err := r.grpcUsersClient.GetUserData(c.Request.Context(), grpcUserDataRequest)
	if err != nil {
		// TODO: check error from gRPC server and return correct error

		server.ErrorResponse[UserFollowCardOutput](c, err.Error())
		return
	}

	var avatarFileKey *string
	if grpcUserDataResponse.GetAvatarFileKey() != nil {
		avatarFileKeyValue := grpcUserDataResponse.GetAvatarFileKey().GetValue()
		avatarFileKey = &avatarFileKeyValue
	}

	userFollowCard := UserFollowCardOutput{
		ID:            grpcUserDataResponse.GetId(),
		FullName:      grpcUserDataResponse.GetFullName(),
		AvatarFileKey: avatarFileKey,
	}

	server.SuccessResponse(c, &userFollowCard)
}

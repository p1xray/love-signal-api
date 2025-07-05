package users

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	lsuserspb "github.com/p1xray/love-signal-protos/gen/go/users"
	urlshortenerpb "github.com/p1xray/pxr-url-shortener/pkg/grpc/gen/go/urlshortener"
	"love-signal-api/internal/controller/http/middleware"
	"love-signal-api/internal/server"
	"time"
)

// Routes provides routes for users.
type Routes struct {
	grpcUsersClient        lsuserspb.UsersClient
	grpcUrlShortenerClient urlshortenerpb.UrlShortenerClient
}

// InitRoutes initializes the routes for users.
func InitRoutes(api *gin.RouterGroup, grpcUsersClient lsuserspb.UsersClient) {
	r := &Routes{grpcUsersClient: grpcUsersClient}

	users := api.Group("/users")
	users.Use(middleware.CheckJWT())
	{
		users.GET("/user-info", middleware.HasScope("users.read"), r.userInfo)
		users.GET("/profile", middleware.HasScope("users.read"), r.userProfileCard)
		users.GET("/follow-link", middleware.HasScope("users.read"), r.followLinkCard)
		users.GET("/follow/:id", middleware.HasScope("users.read"), r.followCard)
		users.GET("/followed", middleware.HasScope("users.read"), r.followedList)
		users.POST("/follow", middleware.HasScope("users.add"), r.follow)
		users.POST("/unfollow", middleware.HasScope("users.delete"), r.unfollow)
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
	userExternalID, err := server.GetUserID(c)
	if err != nil {
		server.ErrorResponse[UserInfoOutput](c, err.Error())
		return
	}

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
	userExternalID, err := server.GetUserID(c)
	if err != nil {
		server.ErrorResponse[UserProfileCardOutput](c, err.Error())
		return
	}

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
	userExternalID, err := server.GetUserID(c)
	if err != nil {
		server.ErrorResponse[UserFollowLinkCardOutput](c, err.Error())
		return
	}

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

	grpcShortenRequest := &urlshortenerpb.ShortenRequest{LongUrl: followLink}
	grpcShortenResponse, err := r.grpcUrlShortenerClient.Shorten(c.Request.Context(), grpcShortenRequest)
	if err != nil {
		// TODO: check error from gRPC server and return correct error

		server.ErrorResponse[UserFollowLinkCardOutput](c, err.Error())
		return
	}

	shortFollowLink := grpcShortenResponse.GetShortUrl()

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

// Followed users list.
//
//	@Summary		Followed users list
//	@Description	Followed users list
//	@Tags			Users
//	@Id 			followedList
//	@Produce		json
//	@Security 		ApiKeyAuth
//	@Success		200	{object}  server.dataResponse[FollowedUsersOutput]
//	@Failure		500	{object}  server.dataResponse[FollowedUsersOutput]
//	@Router			/api/v1/users/followed [get]
func (r *Routes) followedList(c *gin.Context) {
	userExternalID, err := server.GetUserID(c)
	if err != nil {
		server.ErrorResponse[FollowedUsersOutput](c, err.Error())
		return
	}

	grpcUserDataRequest := &lsuserspb.GetUserDataByExternalIdRequest{UserExternalId: userExternalID}
	grpcUserDataResponse, err := r.grpcUsersClient.GetUserDataByExternalId(c.Request.Context(), grpcUserDataRequest)
	if err != nil {
		// TODO: check error from gRPC server and return correct error

		server.ErrorResponse[FollowedUsersOutput](c, err.Error())
		return
	}

	grpcFollowedUsersRequest := &lsuserspb.GetFollowedUsersRequest{UserId: grpcUserDataResponse.GetId()}
	grpcFollowedUsersResponse, err := r.grpcUsersClient.GetFollowedUsers(c.Request.Context(), grpcFollowedUsersRequest)
	if err != nil {
		// TODO: check error from gRPC server and return correct error

		server.ErrorResponse[FollowedUsersOutput](c, err.Error())
		return
	}

	followedUsers := make([]FollowedUserOutput, 0, len(grpcFollowedUsersResponse.Users))
	for _, user := range grpcFollowedUsersResponse.Users {
		var avatarFileKey *string
		if grpcUserDataResponse.GetAvatarFileKey() != nil {
			avatarFileKeyValue := grpcUserDataResponse.GetAvatarFileKey().GetValue()
			avatarFileKey = &avatarFileKeyValue
		}

		followedUser := FollowedUserOutput{
			FollowLinkID:  user.GetFollowLinkId(),
			UserID:        user.GetUserId(),
			FullName:      user.GetFullName(),
			AvatarFileKey: avatarFileKey,
			NumberOfLikes: user.GetNumberOfLikes(),
		}
		followedUsers = append(followedUsers, followedUser)
	}

	output := FollowedUsersOutput{Users: followedUsers}

	server.SuccessResponse(c, &output)
}

// Follow user.
//
//	@Summary		Follow user
//	@Description	Follow user
//	@Tags			Users
//	@Id 			follow
//	@Produce		json
//	@Security 		ApiKeyAuth
//	@Param			input body FollowInput true "Input parameters for follow user."
//	@Success		200	{object}  server.dataResponse[bool]
//	@Failure		500	{object}  server.dataResponse[bool]
//	@Router			/api/v1/users/follow [post]
func (r *Routes) follow(c *gin.Context) {
	userExternalID, err := server.GetUserID(c)
	if err != nil {
		server.ErrorResponse[bool](c, err.Error())
		return
	}

	inp, err := server.GetInputFromBody[FollowInput](c)
	if err != nil {
		server.ErrorResponse[bool](c, err.Error())
		return
	}

	grpcUserDataRequest := &lsuserspb.GetUserDataByExternalIdRequest{UserExternalId: userExternalID}
	grpcUserDataResponse, err := r.grpcUsersClient.GetUserDataByExternalId(c.Request.Context(), grpcUserDataRequest)
	if err != nil {
		// TODO: check error from gRPC server and return correct error

		server.ErrorResponse[bool](c, err.Error())
		return
	}

	grpcFollowUserRequest := &lsuserspb.FollowUserRequest{
		UserId:         grpcUserDataResponse.GetId(),
		UserIdToFollow: inp.UserIDToFollow,
	}
	grpcFollowUserResponse, err := r.grpcUsersClient.FollowUser(c.Request.Context(), grpcFollowUserRequest)
	if err != nil {
		// TODO: check error from gRPC server and return correct error

		server.ErrorResponse[bool](c, err.Error())
		return
	}

	success := grpcFollowUserResponse.GetSuccess()
	server.SuccessResponse[bool](c, &success)
}

// Unfollow user.
//
//	@Summary		Unfollow user
//	@Description	Unfollow user
//	@Tags			Users
//	@Id 			unfollow
//	@Produce		json
//	@Security 		ApiKeyAuth
//	@Param			input body UnfollowInput true "Input parameters for unfollow user."
//	@Success		200	{object}  server.dataResponse[bool]
//	@Failure		500	{object}  server.dataResponse[bool]
//	@Router			/api/v1/users/unfollow [post]
func (r *Routes) unfollow(c *gin.Context) {
	inp, err := server.GetInputFromBody[UnfollowInput](c)
	if err != nil {
		server.ErrorResponse[bool](c, err.Error())
		return
	}

	grpcUnfollowUserRequest := &lsuserspb.UnfollowUserRequest{FollowLinkId: inp.FollowLinkID}
	grpcUnfollowUserResponse, err := r.grpcUsersClient.UnfollowUser(c.Request.Context(), grpcUnfollowUserRequest)
	if err != nil {
		// TODO: check error from gRPC server and return correct error

		server.ErrorResponse[bool](c, err.Error())
		return
	}

	success := grpcUnfollowUserResponse.GetSuccess()
	server.SuccessResponse[bool](c, &success)
}

package users

import "time"

// GenderEnum is type for gender enum.
type GenderEnum int16 // @name GenderEnum

// Gender enum.
const (
	MALE   GenderEnum = 1 // Male
	FEMALE GenderEnum = 2 // Female
)

// UserInfoOutput is output model of user info request.
type UserInfoOutput struct {
	ID       int64  `json:"id"`        // User ID.
	FullName string `json:"full_name"` // User full name.
} // @name UserInfoOutput

// UserProfileCardOutput is output model of user profile card request.
type UserProfileCardOutput struct {
	ID            int64       `json:"id"`              // User ID.
	FullName      string      `json:"full_name"`       // User full name.
	Gender        *GenderEnum `json:"gender"`          // User gender.
	DateOfBirth   *time.Time  `json:"date_of_birth"`   // User date of birth.
	AvatarFileKey *string     `json:"avatar_file_key"` // User avatar file key.
} // @name UserProfileCardOutput

// UserFollowLinkCardOutput is output model of user follow link card request.
type UserFollowLinkCardOutput struct {
	ShortLink string `json:"short_link"` // User follow short link.
	QRCode    string `json:"qr_code"`    // QR code of user follow short link as base64.
} // @name UserFollowLinkCardOutput

// UserFollowCardOutput is output model of user follow card request.
type UserFollowCardOutput struct {
	ID            int64   `json:"id"`              // User ID.
	FullName      string  `json:"full_name"`       // User full name.
	AvatarFileKey *string `json:"avatar_file_key"` // User avatar file key.
} // @name UserFollowCardOutput

// FollowedUsersOutput is output model of followed users list request.
type FollowedUsersOutput struct {
	Users []FollowedUserOutput `json:"users"` // Followed users list.
} // @name FollowedUsersOutput

// FollowedUserOutput is output model of followed user for list.
type FollowedUserOutput struct {
	FollowLinkID  int64   `json:"follow_link_id"`  // Follow link ID.
	UserID        int64   `json:"user_id"`         // User ID.
	FullName      string  `json:"full_name"`       // User full name.
	AvatarFileKey *string `json:"avatar_file_key"` // User avatar file key.
	NumberOfLikes uint32  `json:"number_of_likes"` // Number of likes.
} // @name FollowedUserOutput

// FollowInput is input model of follow user request.
type FollowInput struct {
	UserIDToFollow int64 `json:"user_id_to_follow"` // User ID to follow.
} // @name FollowInput

// UnfollowInput is input model of unfollow user request.
type UnfollowInput struct {
	FollowLinkID int64 `json:"follow_link_id"` // Follow link ID.
} // @name UnfollowInput

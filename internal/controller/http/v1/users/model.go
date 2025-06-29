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

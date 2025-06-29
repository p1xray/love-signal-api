package users

// UserInfoOutput is output model of user info request.
type UserInfoOutput struct {
	ID       int64  `json:"id"`        // User ID.
	FullName string `json:"full_name"` // User full name.
} // @name UserInfoOutput

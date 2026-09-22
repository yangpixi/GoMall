package profile

// UpdateProfileCommand use pointer type to distinguish
// if user change the column.
// this command don't process avatar changing
type UpdateProfileCommand struct {
	UserID   int64
	Nickname *string
	Phone    *string
	Email    *string
}

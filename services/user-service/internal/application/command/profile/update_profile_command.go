package profile

// UpdateProfileCommand use pointer type to distinguish
// if user change the column.
// this command don't process avatar changing
type UpdateProfileCommand struct {
	Nickname *string
	Phone    *string
	Email    *string
}

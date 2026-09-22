package profile

// UpdateProfileCommand use pointer type to distinguish
// if user change the column
type UpdateProfileCommand struct {
	Nickname *string
	Phone    *string
	Email    *string
	Avatar   *string
}

package command

type CreateProfileCommand struct {
	UserID   int64 // only for admin side
	Nickname string
	Phone    string
	Email    string
	Avatar   string
}

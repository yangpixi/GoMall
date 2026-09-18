package command

type CreateProfileCommand struct {
	UserID   int64
	Nickname string
	Phone    string
	Email    string
	Avatar   string
}

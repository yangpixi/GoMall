package command

type TokenIssuer interface {
	Token(username string, roleIDs []uint) (string, int, error)
	RefreshToken(username string, roleIDs []uint) (string, error)
}

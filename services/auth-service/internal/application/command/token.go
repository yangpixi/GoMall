package command

type TokenIssuer interface {
	Token(sub string, roleIDs []int64) (string, int, error)
	RefreshToken(sub string, roleIDs []int64) (string, error)
	ParseAndValidate(token string) (sub, kind string, roleIDs []int64, err error)
}

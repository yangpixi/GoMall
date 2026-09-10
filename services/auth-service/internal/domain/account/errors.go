package account

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrBannedAccount             = errs.New(10001, "account has been banned")
	ErrNotFound                  = errs.New(10002, "account not found")
	ErrUsernamePasswordIncorrect = errs.New(10003, "username or password incorrect")
)

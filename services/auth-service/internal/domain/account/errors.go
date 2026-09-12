package account

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrBannedAccount             = errs.New(10001, "account has been banned")
	ErrAccountNotFound           = errs.New(10002, "account not found")
	ErrUsernamePasswordIncorrect = errs.New(10003, "username or password incorrect")
	ErrMissingUsername           = errs.New(10004, "missing username")
	ErrMissingPassword           = errs.New(10005, "missing password")
	ErrMissingRoles              = errs.New(10006, "missing roles")
)

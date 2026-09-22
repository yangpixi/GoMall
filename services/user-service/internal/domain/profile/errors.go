package profile

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrInvalidProfile   = errs.New(20001, "invalid profile arguments")
	ErrInvalidUserID    = errs.New(20002, "invalid user id")
	ErrInvalidAddressID = errs.New(20003, "invalid address id")
	ErrProfileNotFound  = errs.New(20004, "profile not found")
	ErrInvalidNickname  = errs.New(20005, "invalid nickname")
	ErrInvalidPhone     = errs.New(20006, "invalid phone number")
	ErrInvalidEmail     = errs.New(20007, "invalid email")
	ErrInvalidAvatar    = errs.New(20008, "invalid avatar")
)

package role

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrPermissionNotFound = errs.New(13004, "permissions not found")
)

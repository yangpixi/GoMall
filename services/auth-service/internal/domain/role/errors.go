package role

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrPermissionNotFound = errs.New(12004, "permissions not found")
)

package sku

import "github.com/yangpixi/GoMall/shared/errs"

var (
	ErrInvalidSKU  = errs.New(31001, "invalid sku")
	ErrSKUNotFound = errs.New(31002, "sku not found")
)

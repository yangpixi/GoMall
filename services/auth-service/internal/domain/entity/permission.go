package entity

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Code     string
	Resource string
	Action   string
	Name     string
}

package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	password string
	status   uint
}

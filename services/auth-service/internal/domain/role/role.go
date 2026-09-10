package role

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	code       string
	name       string
	permission []uint
}

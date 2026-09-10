package permission

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	code     string
	resource string
	action   string
	name     string
}

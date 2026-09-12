package model

import "time"

type UserRole struct {
	UserID    uint
	RoleID    uint
	GrantedAt time.Time
}

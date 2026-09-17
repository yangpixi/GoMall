package model

import "time"

type UserRole struct {
	UserID    int64
	RoleID    int64
	GrantedAt time.Time
}

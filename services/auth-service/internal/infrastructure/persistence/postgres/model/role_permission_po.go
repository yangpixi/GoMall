package model

import "time"

type RolePermission struct {
	RoleID       uint
	PermissionID uint
	GrantedAt    time.Time
}

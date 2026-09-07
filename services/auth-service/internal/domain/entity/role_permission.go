package entity

import "time"

type RolePermission struct {
	RoleID       uint
	PermissionID uint
	GrantedAt    time.Time
}

package model

import "time"

type RolePermission struct {
	RoleID       int64
	PermissionID int64
	GrantedAt    time.Time
}

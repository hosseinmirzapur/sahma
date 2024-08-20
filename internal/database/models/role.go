package models

import "time"

type Role struct {
	ID           uint
	Title        string
	Slug         string
	PermissionID uint
	CreatedAt    time.Time
	UpdatedAt    *time.Time
	Permission   *Permission
}

package models

import "time"

type Permission struct {
	ID        uint
	Full      int
	Modify    int
	ReadOnly  int
	CreatedAt time.Time
	UpdatedAt *time.Time
	RoleID    uint
	Role      Role
}

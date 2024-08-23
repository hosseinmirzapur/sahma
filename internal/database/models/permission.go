package models

import "time"

type Permission struct {
	ID        uint
	Full      bool
	Modify    bool
	ReadOnly  bool
	CreatedAt time.Time
	UpdatedAt *time.Time
	RoleID    uint
	Role      Role
}

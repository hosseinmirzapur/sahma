package models

import "time"

type DepartmentFile struct {
	ID            uint
	EntityGroupID uint
	DepartmentID  uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Department    Department
	EntityGroup   EntityGroup
}

package models

import "time"

type DepartmentUser struct {
	ID           uint
	UserID       uint
	DepartmentID uint
	CreatedAt    time.Time
	UpdatedAt    *time.Time
	Department   Department
	User         User
}

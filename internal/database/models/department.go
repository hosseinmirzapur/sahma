package models

import "time"

type Department struct {
	ID              uint
	Name            string
	CreatedBy       *uint
	CreatedAt       time.Time
	UpdatedAt       *time.Time
	DepartmentUsers []*DepartmentUser
	Folders         []*Folder
}

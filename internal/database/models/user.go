package models

import (
	"sahma/internal/globals"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint
	Name            string
	PersonalID      int
	Email           *string
	Password        string
	RoleID          uint
	Meta            *string
	IsSuperAdmin    bool
	RememberToken   *string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
	DeletedAt       *time.Time
	Activities      []*Activity `gorm:"polymorphic:Activity"`
	Letters         []*Letter
	Notifications   []*Notification
	Role            *Role
	Tokens          []*PersonalAccessToken
	UserDepartments []*DepartmentUser
}

func (user *User) GetAllAvailableFiles(folder Folder) ([]map[string]interface{}, error) {
	var egs []EntityGroup
	q, err := user.QueryDepartmentFiles(folder)
	if err != nil {
		return nil, err
	}

	err = q.
		Where("entity_groups.deleted_at = NULL").
		Where("entity_groups.archived_at = NULL").
		Distinct().
		Find(&egs).
		Error
	if err != nil {
		return nil, err
	}
	var results []map[string]interface{}
	for _, eg := range egs {
		slug, err := eg.GetEntityGroupID()
		if err != nil {
			return nil, err
		}
		parentSlug := "" // todo: implement ParentSlug logic

		results = append(results, map[string]interface{}{
			"id":          eg.ID,
			"name":        eg.Name,
			"type":        eg.Type,
			"status":      eg.Status,
			"slug":        *slug,
			"description": *eg.Description,
			"parentSlug":  parentSlug,
		})
	}

	return results, nil
}

func (user *User) QueryDepartmentFiles(parentFolder Folder) (*gorm.DB, error) {
	// Fetch departments
	var departments []DepartmentUser
	err := globals.
		GetDB().
		Where("user_id = ?", user.ID).
		Pluck("department_users.department_id", &departments).
		Error
	if err != nil {
		return nil, err
	}

	// Return query as result
	q := globals.
		GetDB().
		Select("entity_groups.*").
		Joins("JOIN department_files ON department_files.entity_group_id = entity_groups.id").
		Where("entity_groups.parent_folder_id = ?", parentFolder.ID).
		Where("department_files.deparment_id = ?", departments)

	return q, nil
}

func (user *User) GetUserDepartmentIDs() ([]DepartmentUser, error) {
	var depIDs []DepartmentUser

	err := globals.
		GetDB().
		Select("departments.id").
		Joins("JOIN department_users ON department_users.department_id = departments.id").
		Where("department_users.user_id = ?", user.ID).
		Pluck("departments", &depIDs).
		Error
	if err != nil {
		return nil, err
	}

	return depIDs, nil
}

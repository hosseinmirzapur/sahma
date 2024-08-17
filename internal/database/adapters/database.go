package adapters

import (
	"sahma/internal/database/models"
	"sahma/internal/globals"
)

func RegisterMysql() error {
	db, err := RegisterMySQL()
	if err != nil {
		return err
	}

	globals.SetDB(db)
	return nil
}

func Migrate() error {
	err := globals.GetDB().AutoMigrate(
		&models.Activity{},
		&models.Department{},
		&models.DepartmentFile{},
		&models.DepartmentUser{},
		&models.Entity{},
		&models.EntityGroup{},
		&models.Folder{},
		&models.Letter{},
		&models.LetterAttachment{},
		&models.LetterInbox{},
		&models.LetterReply{},
		&models.LetterSign{},
		&models.Notification{},
		&models.Permission{},
		&models.Role{},
		&models.User{},
	)
	if err != nil {
		return err
	}

	// if no error happened then return nil
	return nil
}

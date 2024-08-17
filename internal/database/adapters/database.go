package adapters

import (
	"sahma/internal/database/global"
	"sahma/internal/database/models"
)

func RegisterMysql() error {
	db, err := RegisterMySQL()
	if err != nil {
		return err
	}

	global.DB = db
	return nil
}

func Migrate() error {
	err := global.DB.AutoMigrate(
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

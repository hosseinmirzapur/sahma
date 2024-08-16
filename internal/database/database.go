package database

import (
	"fmt"
	"os"
	"sahma/internal/database/models"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db_user = os.Getenv("DB_USERNAME")
	db_pass = os.Getenv("DB_PASSWORD")
	db_host = os.Getenv("DB_HOST")
	db_port = os.Getenv("DB_PORT")
	db_name = os.Getenv("DB_DATABASE")

	DB *gorm.DB
)

func Register() error {
	// Try to open connection to mysql database
	db, err := gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
			db_user,
			db_pass,
			db_host,
			db_port,
			db_name,
		),
	))
	if err != nil {
		return err
	}

	// Verify db connection
	mysqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if err = mysqlDB.Ping(); err != nil {
		return err
	}

	// when connection is successfull,
	// DB is set globally available for database operations
	DB = db

	return nil

}

func Migrate() error {
	err := DB.AutoMigrate(
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

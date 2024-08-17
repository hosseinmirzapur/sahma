package adapters

import (
	"fmt"
	"os"

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
)

func RegisterMySQL() (*gorm.DB, error) {
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
		return nil, err
	}

	// Verify db connection
	mysqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err = mysqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}

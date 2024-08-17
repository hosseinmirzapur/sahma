package globals

import "gorm.io/gorm"

var globalDB *gorm.DB

func SetDB(db *gorm.DB) {
	globalDB = db
}

func GetDB() *gorm.DB {
	return globalDB
}

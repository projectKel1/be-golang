package database

import (
	_userData "group-project-3/features/user/data"

	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_userData.User{})
}

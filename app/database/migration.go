package database

import (
	_companyData "group-project-3/features/company/data"
	_employeeLevel "group-project-3/features/employeeLevel/data"
	_roleData "group-project-3/features/role/data"
	_userData "group-project-3/features/user/data"
	_userDetailData "group-project-3/features/userDetail/data"

	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_userData.User{})
	db.AutoMigrate(&_userDetailData.UserDetail{})
	db.AutoMigrate(&_roleData.Role{})
	db.AutoMigrate(&_companyData.Company{})
	db.AutoMigrate(&_employeeLevel.EmployeeLevel{})
}

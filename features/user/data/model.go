package data

import (
	_companyData "group-project-3/features/company/data"
	_levelData "group-project-3/features/employeeLevel/data"
	_roleData "group-project-3/features/role/data"
	"group-project-3/features/user"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey"`
	// UserId    uint
	Fullname  string `gorm:"type:varchar(100)"`
	Email     string `gorm:"unique;size:255"`
	Password  string `gorm:"type:varchar(255);unique_index"`
	RoleID    uint
	Role      _roleData.Role
	CompanyID uint
	Company   _companyData.Company
	ManagerID uint
	Manager   *User
	UrlPhoto  string
	Status    string `gorm:"type:enum('Active','Non-Active');column:status;default:Non-Active"`
	LevelID   uint
	Level     _levelData.EmployeeLevel
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CoreToModel(dataCore user.Core) User {
	return User{
		// ID:               dataCore.,
		Fullname:  dataCore.Fullame,
		Email:     dataCore.Email,
		Password:  dataCore.Password,
		RoleID:    dataCore.RoleID,
		ManagerID: dataCore.ManagerID,
		Role:      _roleData.Role{},
		Company:   _companyData.Company{},
		Level:     _levelData.EmployeeLevel{},
		LevelID:   dataCore.LevelID,
		UrlPhoto:  dataCore.UrlPhoto,
		Status:    dataCore.Status,
		CompanyID: dataCore.CompanyID,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel User) user.Core {
	return user.Core{
		ID:        dataModel.ID,
		Fullame:   dataModel.Fullname,
		RoleID:    dataModel.RoleID,
		ManagerID: dataModel.ManagerID,
		Role:      user.RoleCore{},
		Company:   user.CompanyCore{},
		Level:     user.LevelCore{},
		Status:    dataModel.Status,
		Password:  dataModel.Password,
		UrlPhoto:  dataModel.UrlPhoto,
		Email:     dataModel.Email,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

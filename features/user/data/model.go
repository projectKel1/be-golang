package data

import (
	_companyData "group-project-3/features/company/data"
	_levelData "group-project-3/features/employeeLevel/data"
	_roleData "group-project-3/features/role/data"
	"group-project-3/features/user"
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	ACTIVE     Status = "Active"
	NOT_ACTIVE Status = "Non-Active"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Fullname  string `gorm:"type:varchar(100)"`
	Email     string `gorm:"unique;size:255"`
	Password  string `gorm:"type:varchar(255);unique_index"`
	RoleID    uint
	Role      _roleData.Role
	CompanyID uint
	Company   _companyData.Company
	ManagerID *uint
	Manager   *User
	UrlPhoto  string
	Status    Status
	LevelID   uint
	Level     _levelData.EmployeeLevel
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CoreToModel(dataCore user.Core) User {
	return User{
		// ID:               dataCore.,
		Fullname: dataCore.Fullame,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		RoleID:   dataCore.RoleID,
		Role:     _roleData.Role{},
		Company:  _companyData.Company{},
		Level:    _levelData.EmployeeLevel{},
		LevelID:  dataCore.LevelID,
		// RoleName:        dataCore.RoleName,
		// Status: Status(dataCore.Status),
		// Address:         dataCore.Address,
		// Gender:          Gender(dataCore.Gender),
		// PhoneNumber:     dataCore.PhoneNumber,
		// UrlPhoto:        dataCore.UrlPhoto,
		// NoNik:           dataCore.NoNik,
		// NoKk:            dataCore.NoKK,
		// NoBpjs:          dataCore.NoBpjs,
		// Npwp:            dataCore.Npwp,
		// EmergencyName:   dataCore.EmergencyName,
		// EmergencyStatus: dataCore.EmergencyStatus,
		// EmergencyPhone: dataCore.EmergencyPhone,
		CompanyID: dataCore.CompanyID,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel User) user.Core {
	return user.Core{
		ID:      dataModel.ID,
		Fullame: dataModel.Fullname,
		RoleID:  dataModel.RoleID,
		Role:    user.RoleCore{},
		Company: user.CompanyCore{},
		Level:   user.LevelCore{},
		// Role:    dataModel,
		// RoleName:        dataModel.RoleName,
		// PhoneNumber:     dataModel.PhoneNumber,
		// CompanyId: dataModel.CompanyID,
		// NoNik:           dataModel.NoNik,
		// NoKK:            dataModel.NoKk,
		// NoBpjs:          dataModel.NoBpjs,
		Password: dataModel.Password,
		UrlPhoto: dataModel.UrlPhoto,
		// Status:   string(dataModel.Status),
		Email: dataModel.Email,
		// Address:         dataModel.Address,
		// Gender:          string(dataModel.Gender),
		// EmergencyName:   dataModel.EmergencyName,
		// EmergencyStatus: dataModel.EmergencyStatus,
		// EmergencyPhone: dataModel.EmergencyPhone,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

package data

import (
	"group-project-3/features/user"
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	ACTIVE     Status = "Active"
	NOT_ACTIVE Status = "Non-Active"
)

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type User struct {
	ID              uint `gorm:"primaryKey"`
	Fullname        string
	Email           string `gorm:"unique"`
	Password        string
	RoleID          uint
	Status          Status
	Address         string
	Gender          Gender
	NoNik           string `gorm:"unique"`
	NoKk            string
	NoBpjs          string `gorm:"unique"`
	Npwp            string `gorm:"unique"`
	PhoneNumber     string
	EmergencyName   string
	EmergencyStatus string
	UrlPhoto        string
	EmergencyPhone  string
	CompanyId       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func CoreToModel(dataCore user.Core) User {
	return User{
		// ID:               dataCore.,
		Fullname:        dataCore.Fullame,
		Email:           dataCore.Email,
		Password:        dataCore.Password,
		RoleID:          dataCore.RoleID,
		Status:          Status(dataCore.Status),
		Address:         dataCore.Address,
		Gender:          Gender(dataCore.Gender),
		PhoneNumber:     dataCore.PhoneNumber,
		UrlPhoto:        dataCore.UrlPhoto,
		NoNik:           dataCore.NoNik,
		NoKk:            dataCore.NoKK,
		NoBpjs:          dataCore.NoBpjs,
		Npwp:            dataCore.Npwp,
		EmergencyName:   dataCore.EmergencyName,
		EmergencyStatus: dataCore.EmergencyStatus,
		EmergencyPhone:  dataCore.EmergencyPhone,
		CompanyId:       dataCore.CompanyId,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       gorm.DeletedAt{},
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel User) user.Core {
	return user.Core{
		ID:              dataModel.ID,
		Fullame:         dataModel.Fullname,
		RoleID:          dataModel.RoleID,
		PhoneNumber:     dataModel.PhoneNumber,
		CompanyId:       dataModel.CompanyId,
		NoNik:           dataModel.NoNik,
		NoKK:            dataModel.NoKk,
		NoBpjs:          dataModel.NoBpjs,
		Password:        dataModel.Password,
		UrlPhoto:        dataModel.UrlPhoto,
		Status:          string(dataModel.Status),
		Email:           dataModel.Email,
		Address:         dataModel.Address,
		Gender:          string(dataModel.Gender),
		EmergencyName:   dataModel.EmergencyName,
		EmergencyStatus: dataModel.EmergencyStatus,
		EmergencyPhone:  dataModel.EmergencyPhone,
		CreatedAt:       dataModel.CreatedAt,
		UpdatedAt:       dataModel.UpdatedAt,
	}
}

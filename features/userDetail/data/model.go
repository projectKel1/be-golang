package data

import (
	// _userData "group-project-3/features/user/data"
	"time"

	"gorm.io/gorm"
)

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type UserDetail struct {
	ID     uint
	UserID uint
	// User            _userData.User
	Gender          Gender
	Address         string `gorm:"type:varchar(55)"`
	Nik             string `gorm:"type:varchar(255);unique"`
	NoKk            string `gorm:"type:varchar(255)"`
	NoBpjs          string `gorm:"unique"`
	Npwp            string `gorm:"unique"`
	PhoneNumber     string `gorm:"type:varchar(15)"`
	EmergencyName   string `gorm:"type:varchar(255)"`
	EmergencyStatus string `gorm:"type:varchar(255)"`
	EmergencyPhone  string `gorm:"type:varchar(15)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

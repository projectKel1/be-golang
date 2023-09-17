package data

import (
	"group-project-3/features/userDetail"
	"time"

	"gorm.io/gorm"
)

type UserDetail struct {
	ID              uint   `gorm:"primaryKey"`
	UserID          uint   `gorm:"unique;not null;column:user_id"`
	Gender          string `gorm:"type:enum('M','F');column:gender;default:M"`
	Address         string `gorm:"column:address"`
	Nik             string `gorm:"column:nik"`
	NoKK            string `gorm:"column:no_kk"`
	NoBPJS          string `gorm:"column:no_bpjs"`
	Npwp            string `gorm:"column:npwp"`
	PhoneNumber     string `gorm:"column:phone_number"`
	EmergencyName   string `gorm:"column:emergency_name"`
	EmergencyStatus string `gorm:"column:emergency_status"`
	EmergencyPhone  string `gorm:"column:emergency_phone"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func CoreToModel(dataCore userDetail.UserDetailEntity) UserDetail {
	return UserDetail{
		ID:              dataCore.ID,
		UserID:          dataCore.UserID,
		Gender:          dataCore.Gender,
		Address:         dataCore.Address,
		Nik:             dataCore.Nik,
		NoKK:            dataCore.NoKK,
		NoBPJS:          dataCore.NoBPJS,
		Npwp:            dataCore.Npwp,
		PhoneNumber:     dataCore.PhoneNumber,
		EmergencyName:   dataCore.EmergencyName,
		EmergencyStatus: dataCore.EmergencyStatus,
		EmergencyPhone:  dataCore.EmergencyPhone,
		CreatedAt:       dataCore.CreatedAt,
		UpdatedAt:       dataCore.UpdatedAt,
		DeletedAt:       gorm.DeletedAt{},
	}
}

func ModelToCore(dataModel UserDetail) userDetail.UserDetailEntity {
	return userDetail.UserDetailEntity{
		ID:              dataModel.ID,
		UserID:          dataModel.UserID,
		Gender:          dataModel.Gender,
		Address:         dataModel.Address,
		Nik:             dataModel.Nik,
		NoKK:            dataModel.NoKK,
		NoBPJS:          dataModel.NoBPJS,
		Npwp:            dataModel.Npwp,
		PhoneNumber:     dataModel.PhoneNumber,
		EmergencyName:   dataModel.EmergencyName,
		EmergencyStatus: dataModel.EmergencyStatus,
		EmergencyPhone:  dataModel.EmergencyPhone,
		CreatedAt:       dataModel.CreatedAt,
		UpdatedAt:       dataModel.UpdatedAt,
	}
}

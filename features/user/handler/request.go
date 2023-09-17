package handler

import (
	"group-project-3/features/user"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Fullame   string `json:"fullname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	RoleID    uint   `json:"role_id"`
	CompanyID uint   `json:"company_id"`
	ManagerID uint   `json:"manager_id"`
	LevelID   uint   `json:"level_id"`
	UrlPhoto  string `json:"url_photo"`
}

// type UserRequest struct {
// 	Fullame         string `json:"fullname"`
// 	Email           string `json:"email"`
// 	Password        string `json:"password"`
// 	Address         string `json:"address"`
// 	PhoneNumber     string `json:"phone_number"`
// 	RoleID          uint   `json:"role_id"`
// 	Status          string `json:"status"`
// 	Gender          string `json:"gender"`
// 	NoNik           string `json:"no_nik"`
// 	NoKk            string `json:"no_kk"`
// 	NoBpjs          string `json:"no_bpjs"`
// 	Npwp            string `json:"npwp"`
// 	EmergencyName   string `json:"emergency_name"`
// 	EmergencyStatus string `json:"emergency_status"`
// 	EmergencyPhone  string `json:"emergency_phone"`
// 	UrlPhoto        string `json:"url_photo"`
// 	CompanyID       uint   `json:"company_id"`
// }

func RequestToCore(input UserRequest) user.Core {
	return user.Core{
		// ID:              0,
		Fullame:   input.Fullame,
		Email:     input.Email,
		Password:  input.Password,
		RoleID:    input.RoleID,
		CompanyID: input.CompanyID,
		ManagerID: input.ManagerID,
		UrlPhoto:  input.UrlPhoto,
		LevelID:   input.LevelID,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

// func RequestToCore(input UserRequest) user.Core {
// 	return user.Core{
// 		// ID:              0,
// 		Fullame:         input.Fullame,
// 		Email:           input.Email,
// 		Password:        input.Password,
// 		RoleID:          input.RoleID,
// 		Status:          input.Status,
// 		PhoneNumber:     input.PhoneNumber,
// 		Address:         input.Address,
// 		Gender:          input.Gender,
// 		NoNik:           input.NoNik,
// 		NoKK:            input.NoKk,
// 		NoBpjs:          input.NoBpjs,
// 		EmergencyName:   input.EmergencyName,
// 		EmergencyStatus: input.EmergencyName,
// 		EmergencyPhone:  input.EmergencyPhone,
// 		UrlPhoto:        input.UrlPhoto,
// 		Npwp:            input.Npwp,
// 		CompanyId:       input.CompanyID,
// 		CreatedAt:       time.Time{},
// 		UpdatedAt:       time.Time{},
// 	}
// }

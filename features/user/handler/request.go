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

type UpdateProfileRequest struct {
	UserId          uint      `json:"user_id"`
	Fullame         string    `json:"fullname"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	Address         string    `json:"address"`
	Gender          string    `json:"gender"`
	UrlPhoto        string    `json:"url_photo"`
	Nik             string    `json:"nik"`
	NoKK            string    `json:"no_kk"`
	NoBPJS          string    `json:"no_bpjs"`
	Npwp            string    `json:"no_npwp"`
	PhoneNumber     string    `json:"phone_number"`
	EmergencyName   string    `json:"emergency_name"`
	EmergencyStatus string    `json:"emergency_status"`
	EmergencyPhone  string    `json:"emergency_phone"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

func RequestToCore(input UserRequest) user.Core {
	return user.Core{
		// ID:              0,
		Fullame:   input.Fullame,
		Email:     input.Email,
		Password:  input.Password,
		RoleID:    input.RoleID,
		CompanyID: input.CompanyID,
		ManagerID: input.ManagerID,
		LevelID:   input.LevelID,
		UrlPhoto:  input.UrlPhoto,

		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func RequestUpdateProfileToCore(input UpdateProfileRequest) user.UserDetailEntity {
	return user.UserDetailEntity{
		// ID:              inpu,
		UserID:          input.UserId,
		Fullame:         input.Fullame,
		Email:           input.Email,
		Password:        input.Password,
		Address:         input.Address,
		Gender:          input.Gender,
		UrlPhoto:        input.UrlPhoto,
		Nik:             input.Nik,
		NoKK:            input.NoKK,
		NoBPJS:          input.NoBPJS,
		Npwp:            input.Npwp,
		PhoneNumber:     input.PhoneNumber,
		EmergencyName:   input.EmergencyName,
		EmergencyStatus: input.EmergencyStatus,
		EmergencyPhone:  input.EmergencyPhone,
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       time.Time{},
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

package handler

import (
	"group-project-3/features/userDetail"
	"time"
)

type UserDetailResponse struct {
	Id              uint      `json:"id"`
	UserID          uint      `json:"user_id"`
	Gender          string    `json:"gender"`
	Address         string    `json:"address"`
	Nik             string    `json:"nik"`
	NoKK            string    `json:"no_kk"`
	NoBPJS          string    `json:"no_bpjs"`
	Npwp            string    `json:"npwp"`
	PhoneNumber     string    `json:"phone_number"`
	EmergencyName   string    `json:"emergency_name"`
	EmergencyStatus string    `json:"emergency_status"`
	EmergencyPhone  string    `json:"emergency_phone"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func UserDetailToResponse(userDetail userDetail.UserDetailEntity) UserDetailResponse {
	return UserDetailResponse{
		Id:              userDetail.ID,
		UserID:          userDetail.UserID,
		Gender:          userDetail.Gender,
		Address:         userDetail.Address,
		Nik:             userDetail.Nik,
		NoKK:            userDetail.NoKK,
		NoBPJS:          userDetail.NoBPJS,
		Npwp:            userDetail.Npwp,
		PhoneNumber:     userDetail.PhoneNumber,
		EmergencyName:   userDetail.EmergencyName,
		EmergencyStatus: userDetail.EmergencyStatus,
		EmergencyPhone:  userDetail.EmergencyPhone,
		CreatedAt:       userDetail.CreatedAt,
		UpdatedAt:       userDetail.UpdatedAt,
	}
}

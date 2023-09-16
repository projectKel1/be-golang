package handler

import (
	"group-project-3/features/userDetail"
)

type UserDetailRequest struct {
	UserID          uint   `json:"user_id"`
	Gender          string `json:"gender"`
	Address         string `json:"address"`
	Nik             string `json:"nik"`
	NoKK            string `json:"no_kk"`
	NoBPJS          string `json:"no_bpjs"`
	Npwp            string `json:"npwp"`
	PhoneNumber     string `json:"phone_number"`
	EmergencyName   string `json:"emergency_name"`
	EmergencyStatus string `json:"emergency_status"`
	EmergencyPhone  string `json:"emergency_phone"`
}

func RequestToEntity(req UserDetailRequest) userDetail.UserDetailEntity {
	return userDetail.UserDetailEntity{
		UserID:          req.UserID,
		Gender:          req.Gender,
		Address:         req.Address,
		Nik:             req.Nik,
		NoKK:            req.NoKK,
		NoBPJS:          req.NoBPJS,
		Npwp:            req.Npwp,
		PhoneNumber:     req.PhoneNumber,
		EmergencyName:   req.EmergencyName,
		EmergencyStatus: req.EmergencyStatus,
		EmergencyPhone:  req.EmergencyPhone,
	}
}

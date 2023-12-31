package handler

import "time"

type UserResponse struct {
	ID          uint   `json:"id"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	UrlPhoto    string `json:"url_photo"`
	Status      string `json:"status"`
	RoleName    string `json:"role_name"`
	LevelName   string `json:"level_name"`
	CompanyName string `json:"company_name"`
	// Level     string    `json:"level"`
	// Address   string    `json:"address"`
	// Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginResponse struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	RoleName    string `json:"role_name"`
	CompanyName string `json:"company_name"`
	Level       string `json:"level"`
	Fullname    string `json:"fullname"`
	CompanyID   uint   `json:"company_id"`
	Token       string `json:"token"`
}

type ProfileResponse struct {
	ID          uint   `json:"id"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Status      string `json:"status"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	RoleName    string `json:"role_name"`
	CompanyName string `json:"company_name"`
	Level       string `json:"level"`

	// RoleID          int       `json:"role_id"`
	// CompanyID int    `json:"company_id"`
	UrlPhoto string `json:"url_photo"`

	NoNik           string    `json:"no_nik"`
	NoKK            string    `json:"no_nokk"`
	NoBpjs          string    `json:"no_bpjs"`
	Npwp            string    `json:"no_npwp"`
	EmergencyName   string    `json:"emergency_name"`
	EmergencyStatus string    `json:"emergency_status"`
	EmergencyPhone  string    `json:"emergency_phone"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type DetailResponse struct {
	ID          uint   `json:"id"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	RoleName    string `json:"role_name"`
	CompanyName string `json:"company_name"`
	Level       string `json:"level"`

	// RoleID          int       `json:"role_id"`
	// CompanyID int    `json:"company_id"`
	UrlPhoto string `json:"url_photo"`

	NoNik           string    `json:"no_nik"`
	NoKK            string    `json:"no_nokk"`
	NoBpjs          string    `json:"no_bpjs"`
	Npwp            string    `json:"no_npwp"`
	EmergencyName   string    `json:"emergency_name"`
	EmergencyStatus string    `json:"emergency_status"`
	EmergencyPhone  string    `json:"emergency_phone"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

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

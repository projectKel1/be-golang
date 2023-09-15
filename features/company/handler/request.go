package handler

import (
	"group-project-3/features/company"
	"time"
)

type CompanyRequest struct {
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	Email       string    `json:"email"`
	Type        string    `json:"type"`
	Image       string    `json:"image"`
	StartedHour time.Time `json:"started_hour"`
	EndedHour   time.Time `json:"ended_hour"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func RequestToCore(input CompanyRequest) company.Core {
	return company.Core{
		Name:        input.Name,
		Address:     input.Address,
		Description: input.Description,
		Email:       input.Email,
		Type:        input.Type,
		Image:       input.Image,
		StartedHour: input.StartedHour,
		EndedHour:   input.EndedHour,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   time.Time{},
	}
}

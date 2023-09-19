package data

import (
	"group-project-3/features/company"
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Address     string
	Description string
	Email       string
	Type        string
	Image       string
	StartedHour string
	EndedHour   string
	Visi        string
	Misi        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func CoreToModel(dataCore company.Core) Company {
	return Company{
		ID:          dataCore.ID,
		Name:        dataCore.Name,
		Address:     dataCore.Address,
		Description: dataCore.Description,
		Email:       dataCore.Email,
		Type:        dataCore.Type,
		Image:       dataCore.Image,
		StartedHour: dataCore.StartedHour,
		EndedHour:   dataCore.EndedHour,
		Visi:        dataCore.Visi,
		Misi:        dataCore.Misi,
		CreatedAt:   dataCore.CreatedAt,
		UpdatedAt:   dataCore.UpdatedAt,
		DeletedAt:   gorm.DeletedAt{},
	}
}

func ModelToCore(dataModel Company) company.Core {
	return company.Core{
		ID:          dataModel.ID,
		Name:        dataModel.Name,
		Address:     dataModel.Address,
		Description: dataModel.Description,
		Email:       dataModel.Email,
		Type:        dataModel.Type,
		Image:       dataModel.Image,
		StartedHour: dataModel.StartedHour,
		EndedHour:   dataModel.EndedHour,
		Visi:        dataModel.Visi,
		Misi:        dataModel.Misi,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
		DeletedAt:   time.Time{},
	}
}

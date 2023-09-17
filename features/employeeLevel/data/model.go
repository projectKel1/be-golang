package data

import (
	"group-project-3/features/employeeLevel"

	"time"

	"gorm.io/gorm"
)

type EmployeeLevel struct {
	ID        uint   `gorm:"primaryKey"`
	Level     string `gorm:"type:varchar(55)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CoreToModel(dataCore employeeLevel.Core) EmployeeLevel {
	return EmployeeLevel{
		ID:        0,
		Level:     dataCore.Level,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel EmployeeLevel) employeeLevel.Core {
	return employeeLevel.Core{
		ID:    0,
		Level: dataModel.Level,
	}
}

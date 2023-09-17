package data

import (
	"group-project-3/features/employeeLevel"
	"time"

	"gorm.io/gorm"
)

type EmployeeLevel struct {
	ID        uint   `gorm:"primaryKey"`
	Level     string `gorm:"column:level;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CoreToModel(dataCore employeeLevel.EmployeeLevelEntity) EmployeeLevel {
	return EmployeeLevel{
		ID:        dataCore.ID,
		Level:     dataCore.Level,
		CreatedAt: dataCore.CreatedAt,
		UpdatedAt: dataCore.UpdatedAt,
		DeletedAt: gorm.DeletedAt{},
	}
}

func ModelToCore(dataModel EmployeeLevel) employeeLevel.EmployeeLevelEntity {
	return employeeLevel.EmployeeLevelEntity{
		ID:        dataModel.ID,
		Level:     dataModel.Level,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

func SliceModelToSliceCore(levels []EmployeeLevel) []employeeLevel.EmployeeLevelEntity {
	var sliceCores []employeeLevel.EmployeeLevelEntity
	for _, model := range levels {
		entity := ModelToCore(model)
		sliceCores = append(sliceCores, entity)
	}
	return sliceCores
}

package handler

import (
	"group-project-3/features/employeeLevel"
)

type EmployeeLevelRequest struct {
	ID    uint   `json:"id"`
	Level string `json:"level"`
}

func RequestToEntity(req EmployeeLevelRequest) employeeLevel.EmployeeLevelEntity {
	return employeeLevel.EmployeeLevelEntity{
		ID:    req.ID,
		Level: req.Level,
	}
}

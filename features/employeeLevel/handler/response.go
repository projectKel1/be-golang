package handler

import (
	"group-project-3/features/employeeLevel"
	"time"
)

type EmployeeLevelResponse struct {
	Id        uint      `json:"id"`
	Level     string    `json:"level"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func EmployeeLevelToResponse(employeeLevel employeeLevel.EmployeeLevelEntity) EmployeeLevelResponse {
	return EmployeeLevelResponse{
		Id:        employeeLevel.ID,
		Level:     employeeLevel.Level,
		CreatedAt: employeeLevel.CreatedAt,
		UpdatedAt: employeeLevel.UpdatedAt,
	}
}

func SliceEmployeeLevelToSliceResponse(employeeLevels []employeeLevel.EmployeeLevelEntity) []EmployeeLevelResponse {
	var SliceResponses []EmployeeLevelResponse
	for _, entity := range employeeLevels {
		response := EmployeeLevelToResponse(entity)
		SliceResponses = append(SliceResponses, response)
	}
	return SliceResponses
}

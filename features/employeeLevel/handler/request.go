package handler

import (
	"group-project-3/features/employeeLevel"
)

type EmployeeLevelRequest struct {
	Level string `json:"level"`
}

func RequestToCore(input EmployeeLevelRequest) employeeLevel.Core {
	return employeeLevel.Core{
		Level: input.Level,
	}
}

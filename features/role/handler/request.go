package handler

import (
	"group-project-3/features/role"
)

type RoleRequest struct {
	RoleNme string `json:"role_name"`
}

func RequestToCore(input RoleRequest) role.Core {
	return role.Core{
		// ID:       0,
		RoleName: input.RoleNme,
	}
}

package data

import (
	"group-project-3/features/role"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint `gorm:"primaryKey"`
	RoleName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CoreToModel(dataCore role.Core) Role {
	return Role{
		ID:        0,
		RoleName:  dataCore.RoleName,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}

// mapping struct model to struct core
func ModelToCore(dataModel Role) role.Core {
	return role.Core{
		ID:       0,
		RoleName: dataModel.RoleName,
	}
}

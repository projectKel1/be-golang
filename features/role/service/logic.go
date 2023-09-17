package service

import (
	"group-project-3/features/role"

	"github.com/go-playground/validator/v10"
)

type roleService struct {
	roleData role.RoleDataInterface
	validate *validator.Validate
}

func New(repo role.RoleDataInterface) role.RoleServiceInterface {
	return &roleService{
		roleData: repo,
	}
}

// Create implements role.RoleServiceInterface.
func (service *roleService) Create(input role.Core) error {
	// errValidate := service.validate.Struct(input)
	// if errValidate != nil {
	// 	return errors.New("validation error" + errValidate.Error())
	// }

	err := service.roleData.Insert(input)
	return err
}

// GetAll implements role.RoleServiceInterface.
func (service *roleService) GetAll() ([]role.Core, error) {
	return service.roleData.SelectAll()
}

// UpdateRole implements role.RoleServiceInterface.
func (service *roleService) UpdateRole(idRole uint, input role.Core) error {
	err := service.roleData.EditRole(idRole, input)
	return err
}

// DeleteRole implements role.RoleServiceInterface.
func (service *roleService) DeleteRole(idRole int) error {
	err := service.roleData.DeleteRole(idRole)
	return err
}

// GetById implements role.RoleServiceInterface.
func (service *roleService) GetById(id uint) (role.Core, error) {
	return service.roleData.SelectById(id)
}

package service

import (
	"group-project-3/features/employeeLevel"

	"github.com/go-playground/validator/v10"
)

type employeeLevelService struct {
	employeeLevelData employeeLevel.EmployeeLevelDataInterface
	validate          *validator.Validate
}

func New(repo employeeLevel.EmployeeLevelDataInterface) employeeLevel.EmployeeLevelServiceInterface {
	return &employeeLevelService{
		employeeLevelData: repo,
	}
}

// Create implements employeelevel.EmployeeLevelServiceInterface.
func (service *employeeLevelService) Create(input employeeLevel.Core) error {
	// errValidate := service.validate.Struct(input)
	// if errValidate != nil {
	// 	return errors.New("validation error" + errValidate.Error())
	// }

	err := service.employeeLevelData.Insert(input)
	return err
}

// GetAll implements employeeLevel.EmployeeLevelServiceInterface.
func (service *employeeLevelService) GetAll() ([]employeeLevel.Core, error) {
	return service.employeeLevelData.SelectAll()
}

// UpdateEmployeeLevel implements employeeLevel.EmployeeLevelServiceInterface.
func (service *employeeLevelService) UpdateEmployeeLevel(idEmployeeLevel uint, input employeeLevel.Core) error {
	err := service.employeeLevelData.EditEmployeeLevel(idEmployeeLevel, input)
	return err
}

// DeleteEmployeeLevel implements employeeLevel.EmployeeLevelServiceInterface.
func (service *employeeLevelService) DeleteEmployeeLevel(idEmployeeLevel int) error {
	err := service.employeeLevelData.DeleteEmployeeLevel(idEmployeeLevel)
	return err
}

package service

import (
	"group-project-3/features/employeeLevel"
	"strings"
)

type EmployeeLevelServiceImpl struct {
	Data employeeLevel.EmployeeLevelDataInterface
}

func New(data employeeLevel.EmployeeLevelDataInterface) employeeLevel.EmployeeLevelServiceInterface {
	return &EmployeeLevelServiceImpl{Data: data}
}

// FindAll implements employeeLevel.EmployeeLevelServiceInterface
func (service *EmployeeLevelServiceImpl) FindAll() ([]employeeLevel.EmployeeLevelEntity, error) {
	results, err := service.Data.FindAll()
	return results, err
}

// Create implements employeeLevel.EmployeeLevelerviceInterface
func (service *EmployeeLevelServiceImpl) Create(employeeLevelEntity employeeLevel.EmployeeLevelEntity) error {
	upperString := strings.ToUpper(employeeLevelEntity.Level)
	employeeLevelEntity.Level = upperString

	err := service.Data.Save(employeeLevelEntity)
	return err
}

// Delete implements employeeLevel.EmployeeLevelerviceInterface
func (service *EmployeeLevelServiceImpl) Delete(id uint) error {
	err := service.Data.Delete(id)
	return err
}

// FindById implements employeeLevel.EmployeeLevelerviceInterface
func (service *EmployeeLevelServiceImpl) FindById(id uint) (employeeLevel.EmployeeLevelEntity, error) {
	result, err := service.Data.FindById(id)
	return result, err
}

// Update implements employeeLevel.EmployeeLevelerviceInterface
func (service *EmployeeLevelServiceImpl) Update(employeeLevelEntity employeeLevel.EmployeeLevelEntity) error {
	upperString := strings.ToUpper(employeeLevelEntity.Level)
	employeeLevelEntity.Level = upperString

	err := service.Data.Update(employeeLevelEntity)
	return err
}

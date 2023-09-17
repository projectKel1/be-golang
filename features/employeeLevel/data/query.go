package data

import (
	"errors"
	"group-project-3/exception"
	"group-project-3/features/employeeLevel"
	"log"

	"gorm.io/gorm"
)

type EmployeeLevelDataImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) employeeLevel.EmployeeLevelDataInterface {
	return &EmployeeLevelDataImpl{db: db}
}

// Delete implements employeeLevel.EmployeeLevelDataInterface
func (data *EmployeeLevelDataImpl) Delete(id uint) error {
	tx := data.db.Delete(&EmployeeLevel{}, id)

	if tx.Error != nil {
		return exception.ErrInternalServer
	}
	if tx.RowsAffected == 0 {
		return exception.ErrIdIsNotFound
	}

	return nil

}

// FindById implements employeeLevel.EmployeeLevelDataInterface
func (data *EmployeeLevelDataImpl) FindById(id uint) (employeeLevel.EmployeeLevelEntity, error) {
	var dataGorm EmployeeLevel
	tx := data.db.First(&dataGorm, id)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return employeeLevel.EmployeeLevelEntity{}, exception.ErrIdIsNotFound
		} else {
			return employeeLevel.EmployeeLevelEntity{}, exception.ErrInternalServer
		}
	}
	employeeLevelEntity := ModelToCore(dataGorm)
	return employeeLevelEntity, nil
}

// Save implements employeeLevel.EmployeeLevelDataInterface
func (data *EmployeeLevelDataImpl) Save(employeeLevel employeeLevel.EmployeeLevelEntity) error {
	employeeLevelGorm := CoreToModel(employeeLevel)
	tx := data.db.Create(&employeeLevelGorm)
	if tx.Error != nil {
		log.Println(tx.Error)
		return exception.ErrInternalServer
	}
	return nil
}

// Update implements employeeLevel.EmployeeLevelDataInterface
func (data *EmployeeLevelDataImpl) Update(employeeLevel employeeLevel.EmployeeLevelEntity) error {
	var employeeLevelGorm EmployeeLevel

	tx := data.db.First(&employeeLevelGorm, employeeLevel.ID)
	if tx.Error != nil {
		return exception.ErrInternalServer
	}

	if tx.RowsAffected == 0 {
		return exception.ErrIdIsNotFound
	}
	userUpdate := CoreToModel(employeeLevel)
	tx = data.db.Model(&employeeLevelGorm).Updates(userUpdate)
	if tx.Error != nil {
		return exception.ErrInternalServer
	}

	if tx.RowsAffected == 0 {
		return exception.ErrIdIsNotFound
	}

	return nil
}

// FindAll implements employeeLevel.EmployeeLevelDataInterface
func (data *EmployeeLevelDataImpl) FindAll() ([]employeeLevel.EmployeeLevelEntity, error) {
	var employeeLevels []EmployeeLevel

	tx := data.db.Find(&employeeLevels)
	if tx.Error != nil {
		return []employeeLevel.EmployeeLevelEntity{}, exception.ErrInternalServer
	}

	employeeLevelEntities := SliceModelToSliceCore(employeeLevels)
	return employeeLevelEntities, nil

}

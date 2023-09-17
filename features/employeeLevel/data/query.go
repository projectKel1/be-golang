package data

import (
	"errors"
	"fmt"
	"group-project-3/features/employeeLevel"
	"log"

	"gorm.io/gorm"
)

type employeeLevelQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) employeeLevel.EmployeeLevelDataInterface {
	return &employeeLevelQuery{
		db: db,
	}
}

// Insert implements employeeLevel.EmployeeLevelDataInterface.
func (repo *employeeLevelQuery) Insert(input employeeLevel.Core) error {
	employeeLevelGorm := CoreToModel(input)

	// simpan ke DB
	tx := repo.db.Create(&employeeLevelGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectAll implements employeeLevel.EmployeeLevelDataInterface.
func (repo *employeeLevelQuery) SelectAll() ([]employeeLevel.Core, error) {
	var employeeLevelData []EmployeeLevel
	var employeeLevelCore []employeeLevel.Core

	tx := repo.db.Raw("SELECT*FROM employee_levels").Scan(&employeeLevelData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	for _, value := range employeeLevelData {
		var roleValue = employeeLevel.Core{
			ID:    value.ID,
			Level: value.Level,
		}
		employeeLevelCore = append(employeeLevelCore, roleValue)
	}
	return employeeLevelCore, nil
}

// EditEmployeeLevel implements employeeLevel.EmployeeLevelDataInterface.
func (repo *employeeLevelQuery) EditEmployeeLevel(idEmployeeLevel uint, input employeeLevel.Core) error {
	if input.Level == "" {
		return errors.New("Level name is required")
	}

	tx := repo.db.Exec("UPDATE employee_levels SET level= ? WHERE id=?", input.Level, idEmployeeLevel)
	fmt.Println("Employee Levels ", tx)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// DeleteEmployeeLevel implements employeeLevel.EmployeeLevelDataInterface.
func (repo *employeeLevelQuery) DeleteEmployeeLevel(idEmployeeLevel int) error {
	var employeeLevel EmployeeLevel
	result := repo.db.Raw("SELECT*FROM employee_levels WHERE id=?", idEmployeeLevel).Scan(&employeeLevel)

	fmt.Println("RESULT ROLE=== ", result)
	fmt.Println("RESULT ROLE=== ", employeeLevel)
	if result.Error != nil {
		log.Fatalf("cannot retrieve employee levels: %v\n", result.Error)
	}
	if result.Error != nil {
		log.Fatalf("cannot delete Publisher: %v\n", result.Error)
	}

	//Delete the role record
	result = repo.db.Exec("DELETE FROM employee_levels WHERE id=?", idEmployeeLevel)
	return nil
}

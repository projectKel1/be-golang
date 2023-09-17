package data

import (
	"errors"
	"fmt"
	"group-project-3/features/role"
	"log"

	"gorm.io/gorm"
)

type roleQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) role.RoleDataInterface {
	return &roleQuery{
		db: db,
	}
}

// Insert implements role.RoleDataInterface.
func (repo *roleQuery) Insert(input role.Core) error {
	roleGorm := CoreToModel(input)

	// simpan ke DB
	tx := repo.db.Create(&roleGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectAll implements role.RoleDataInterface.
func (repo *roleQuery) SelectAll() ([]role.Core, error) {
	var roleData []Role
	var roleCore []role.Core

	tx := repo.db.Raw("SELECT*FROM roles").Scan(&roleData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	for _, value := range roleData {
		var roleValue = role.Core{
			ID:       value.ID,
			RoleName: value.RoleName,
		}
		roleCore = append(roleCore, roleValue)
	}
	return roleCore, nil

}

// EditRole implements role.RoleDataInterface.
func (repo *roleQuery) EditRole(idRole uint, input role.Core) error {
	if input.RoleName == "" {
		return errors.New("Role name is required")
	}

	tx := repo.db.Exec("UPDATE roles SET role_name= ? WHERE id=?", input.RoleName, idRole)
	fmt.Println("ROLE ", tx)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// DeleteRole implements role.RoleDataInterface.
func (repo *roleQuery) DeleteRole(idRole int) error {
	var role Role
	result := repo.db.Raw("SELECT*FROM roles WHERE id=?", idRole).Scan(&role)

	fmt.Println("RESULT ROLE=== ", result)
	fmt.Println("RESULT ROLE=== ", role)
	if result.Error != nil {
		log.Fatalf("cannot retrieve data role: %v\n", result.Error)
	}
	if result.Error != nil {
		log.Fatalf("cannot delete Publisher: %v\n", result.Error)
	}

	//Delete the role record
	result = repo.db.Exec("DELETE FROM roles WHERE id=?", idRole)
	return nil
}

// SelectById implements role.RoleDataInterface.
func (repo *roleQuery) SelectById(id uint) (role.Core, error) {
	var result Role
	tx := repo.db.Raw("SELECT *FROM roles WHERE id=?", id).Scan(&result)
	fmt.Println(tx)
	if tx.Error != nil {
		return role.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return role.Core{}, errors.New("data not found")
	}

	resultCore := ModelToCore(result)
	return resultCore, nil
}

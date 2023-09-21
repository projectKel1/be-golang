package data

import (
	"errors"
	"group-project-3/features/company"
	"time"

	"gorm.io/gorm"
)

type companyQuery struct {
	db *gorm.DB
}

// Delete implements company.CompanyDataInterface.
func (repo *companyQuery) Delete(id uint) error {
	var companyGorm Company
	tx := repo.db.Where("id = ?", id).Delete(&companyGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Insert implements company.CompanyDataInterface.
func (repo *companyQuery) Insert(input company.Core) error {

	companyGorm := CoreToModel(input)
	tx := repo.db.Create(&companyGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Select implements company.CompanyDataInterface.
func (repo *companyQuery) Select(id uint) (company.Core, error) {
	var companyData Company
	tx := repo.db.Where("id = ?", id).First(&companyData)
	if tx.Error != nil {
		return company.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return company.Core{}, errors.New("data not found")
	}

	return ModelToCore(companyData), nil
}

// SelectAll implements company.CompanyDataInterface.
func (repo *companyQuery) SelectAll(pageNumber int, pageSize int) ([]company.Core, error) {
	var companyData []Company
	offset := (pageNumber - 1) * pageSize

	tx := repo.db.Offset(offset).Limit(pageSize).Find(&companyData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var companyCore []company.Core
	for _, value := range companyData {
		companyCore = append(companyCore, company.Core{
			ID:          value.ID,
			Name:        value.Name,
			Address:     value.Address,
			Description: value.Description,
			Email:       value.Email,
			Type:        value.Type,
			StartedHour: value.StartedHour,
			EndedHour:   value.EndedHour,
			Visi:        value.Visi,
			Misi:        value.Misi,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeletedAt:   time.Time{},
		})
	}
	return companyCore, nil
}

// Update implements company.CompanyDataInterface.
func (repo *companyQuery) Update(id uint, input company.Core) error {

	companyGorm := CoreToModel(input)
	tx := repo.db.Model(&Company{}).Where("id = ?", id).Updates(&companyGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func New(db *gorm.DB) company.CompanyDataInterface {
	return &companyQuery{
		db: db,
	}
}

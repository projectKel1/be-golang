package service

import (
	"errors"
	"group-project-3/features/company"

	"github.com/go-playground/validator/v10"
)

type companyService struct {
	companyData company.CompanyDataInterface
	validate    *validator.Validate
}

// Create implements company.CompanyServiceInterface.
func (service *companyService) Create(input company.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("validation error" + errValidate.Error())
	}

	err := service.companyData.Insert(input)
	return err
}

// DeleteById implements company.CompanyServiceInterface.
func (service *companyService) DeleteById(id uint) error {
	err := service.companyData.Delete(id)
	return err
}

// EditById implements company.CompanyServiceInterface.
func (service *companyService) EditById(id uint, input company.Core) error {
	err := service.companyData.Update(id, input)
	return err
}

// GetAll implements company.CompanyServiceInterface.
func (service *companyService) GetAll(pageNumber int, pageSize int) ([]company.Core, error) {

	result, err := service.companyData.SelectAll(pageNumber, pageSize)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// GetById implements company.CompanyServiceInterface.
func (service *companyService) GetById(id uint) (company.Core, error) {
	result, err := service.companyData.Select(id)
	if err != nil {
		return company.Core{}, err
	}
	return result, nil
}

func New(repo company.CompanyDataInterface) company.CompanyServiceInterface {
	return &companyService{
		companyData: repo,
		validate:    validator.New(),
	}
}

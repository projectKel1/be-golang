package service

import (
	"errors"
	"group-project-3/app/middlewares"
	"group-project-3/features/user"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.Core, companyId int) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("validation error" + errValidate.Error())
	}

	err := service.userData.Insert(input, companyId)
	return err
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (dataLogin user.Core, token string, err error) {

	dataLogin, err = service.userData.Login(email, password)
	if err != nil {
		return user.Core{}, "", err
	}
	token, err = middlewares.CreateToken(dataLogin.ID, dataLogin.Role.RoleName, dataLogin.Level.Level, dataLogin.Company.CompanyName, dataLogin.Fullame, dataLogin.Company.ID)
	if err != nil {
		return user.Core{}, "", err
	}
	return dataLogin, token, nil
}

// GetProfile implements user.UserServiceInterface.
func (service *userService) GetProfile(id int) (user.Core, error) {
	result, err := service.userData.SelectProfile(id)
	return result, err
}

// GetAll implements user.UserServiceInterface.
func (service *userService) GetAll(pageNumber int, pageSize int, managerId int, companyId int, filterManager int) ([]user.Core, error) {
	result, err := service.userData.SelectAll(pageNumber, pageSize, managerId, companyId, filterManager)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update implements user.UserServiceInterface.
func (service *userService) UpdateProfile(id int, input user.UserDetailEntity) error {
	err := service.userData.UpdateProfile(id, input)
	return err
}

// UpdateOtherProfile implements user.UserServiceInterface.
func (service *userService) UpdateOtherProfile(id int, input user.UserDetailEntity) error {
	err := service.userData.UpdateOtherProfile(id, input)
	return err
}

// SelectOtherProfile implements user.UserServiceInterface.
func (service *userService) SelectOtherProfile(id int) (user.Core, error) {
	result, err := service.userData.SelectOtherProfile(id)
	return result, err
}

// DeleteOtherProfile implements user.UserServiceInterface.
func (service *userService) DeleteOtherProfile(id uint) error {
	err := service.userData.DeleteOtherProfile(id)
	return err
}

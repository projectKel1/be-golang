package service

import (
	"errors"
	"fmt"
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
func (service *userService) Create(input user.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errors.New("validation error" + errValidate.Error())
	}

	err := service.userData.Insert(input)
	return err
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (dataLogin user.Core, token string, err error) {

	dataLogin, err = service.userData.Login(email, password)
	fmt.Println("ISI DATA LOGIN", dataLogin.Role.RoleName)
	fmt.Println("ISI DATA LOGIN", dataLogin.Company.CompanyName)
	fmt.Println("ISI DATA LOGIN", dataLogin.Level.Level)
	if err != nil {
		return user.Core{}, "", err
	}
	token, err = middlewares.CreateToken(dataLogin.ID, dataLogin.Role.RoleName, dataLogin.Level.Level, dataLogin.Company.CompanyName)
	if err != nil {
		return user.Core{}, "", err
	}
	return dataLogin, token, nil
}

// GetProfile implements user.UserServiceInterface.
func (service *userService) GetProfile(id int) (dataProfile user.Core, err error) {
	fmt.Println("ISI DATA PROFILE", dataProfile.Role.RoleName)
	result, err := service.userData.SelectProfile(id)
	return result, err
}

// GetAll implements user.UserServiceInterface.
func (service *userService) GetAll(pageNumber int, pageSize int) ([]user.Core, error) {
	result, err := service.userData.SelectAll(pageNumber, pageSize)

	if err != nil {
		return nil, err
	}
	return result, nil
}

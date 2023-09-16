package userDetail

import (
	"time"

	"github.com/labstack/echo/v4"
)

type UserDetailEntity struct {
	ID              uint
	UserID          uint
	Gender          string
	Address         string
	Nik             string
	NoKK            string
	NoBPJS          string
	Npwp            string
	PhoneNumber     string
	EmergencyName   string
	EmergencyStatus string
	EmergencyPhone  string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type UserDetailDataInterface interface {
	Save(userDetail UserDetailEntity) error
	Update(userDetail UserDetailEntity) error
	Delete(id uint) error
	FindById(id uint) (UserDetailEntity, error)
}

type UserDetailServiceInterface interface {
	Create(userDetail UserDetailEntity) error
	Update(userDetail UserDetailEntity) error
	Delete(id uint) error
	FindById(id uint) (UserDetailEntity, error)
}

type UserDetailHandlerInterface interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	FindById(c echo.Context) error
}

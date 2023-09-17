package employeeLevel

import (
	"time"

	"github.com/labstack/echo/v4"
)

type EmployeeLevelEntity struct {
	ID        uint
	Level     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type EmployeeLevelDataInterface interface {
	Save(employeeLevel EmployeeLevelEntity) error
	Update(EmployeeLevel EmployeeLevelEntity) error
	Delete(id uint) error
	FindById(id uint) (EmployeeLevelEntity, error)
	FindAll() ([]EmployeeLevelEntity, error)
}

type EmployeeLevelServiceInterface interface {
	Create(EmployeeLevel EmployeeLevelEntity) error
	Update(EmployeeLevel EmployeeLevelEntity) error
	Delete(id uint) error
	FindById(id uint) (EmployeeLevelEntity, error)
	FindAll() ([]EmployeeLevelEntity, error)
}

type EmployeeLevelHandlerInterface interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	FindById(c echo.Context) error
	FindAll(c echo.Context) error
}

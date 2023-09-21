package company

import "time"

type Core struct {
	ID          uint
	Name        string `validate:"required"`
	Address     string `validate:"required"`
	Description string
	Email       string `validate:"required,email"`
	Type        string
	StartedHour string
	EndedHour   string
	Visi        string
	Misi        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type CompanyDataInterface interface {
	Insert(input Core) error
	SelectAll(pageNumber int, pageSize int) ([]Core, error)
	Update(id uint, input Core) error
	Delete(id uint) error
	Select(id uint) (Core, error)
}

type CompanyServiceInterface interface {
	Create(input Core) error
	GetAll(pageNumber int, pageSize int) ([]Core, error)
	EditById(id uint, input Core) error
	DeleteById(id uint) error
	GetById(id uint) (Core, error)
}

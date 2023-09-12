package user

import "time"

type Core struct {
	ID              uint
	Fullame         string `validate:"required"`
	Email           string `validate:"required,email"`
	Password        string `validate:"required"`
	RoleID          uint   `validate:"required"`
	Status          string
	PhoneNumber     string
	Address         string
	Gender          string
	NoNik           string
	NoKK            string
	NoBpjs          string
	EmergencyName   string
	EmergencyStatus string
	EmergencyPhone  string
	UrlPhoto        string
	Npwp            string
	CompanyId       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type UserDataInterface interface {
	Insert(input Core) error
	Login(email string, password string) (dataLogin Core, err error)
	SelectProfile(id int) (Core, error)
}

type UserServiceInterface interface {
	Create(input Core) error
	Login(email string, password string) (dataLogin Core, token string, err error)
	GetProfile(id int) (Core, error)
}

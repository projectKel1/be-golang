package user

import (
	"time"
)

type Core struct {
	ID          uint
	Fullame     string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	RoleID      uint   `validate:"required"`
	CompanyID   uint   `validate:"required"`
	ManagerID   uint
	RoleName    string
	LevelName   string
	CompanyName string
	UrlPhoto    string
	LevelID     uint `validate:"required"`
	Role        RoleCore
	Company     CompanyCore
	Level       LevelCore
	UserDetail  UserDetailCore
	// Status      string
	// PhoneNumber string
	// Address     string
	// Gender      string
	// NoNik       string
	// NoKK            string
	// NoBpjs          string
	// EmergencyName   string
	// EmergencyStatus string
	// EmergencyPhone  string

	// Npwp            string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoleCore struct {
	ID       uint
	RoleName string
}

type CompanyCore struct {
	ID          uint
	CompanyName string
}
type LevelCore struct {
	ID    uint
	Level string
}

type UserDetailCore struct {
	ID          uint
	Address     string
	Nik         string
	Gender      string
	PhoneNumber string
	NoKK        string
	Bpjs        string
	Npwp        string
}
type UserDataInterface interface {
	Insert(input Core) error
	Login(email string, password string) (dataLogin Core, err error)
	SelectProfile(id int) (dataProfile Core, err error)
	SelectAll(pageNumber int, pageSize int) ([]Core, error)
}

type UserServiceInterface interface {
	Create(input Core) error
	Login(email string, password string) (dataLogin Core, token string, err error)
	GetProfile(id int) (dataProfile Core, err error)
	GetAll(pageNumber int, pageSize int) ([]Core, error)
}

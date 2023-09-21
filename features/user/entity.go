package user

import (
	"time"
)

type Core struct {
	ID          uint
	Fullame     string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	RoleID      uint
	CompanyID   uint
	ManagerID   uint
	RoleName    string
	Status      string
	LevelName   string
	CompanyName string
	UrlPhoto    string
	LevelID     uint
	Role        RoleCore
	Company     CompanyCore
	Level       LevelCore
	UserDetail  UserDetailCore
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
	ID              uint
	Address         string
	Nik             string
	Gender          string
	PhoneNumber     string
	NoKK            string
	Bpjs            string
	Npwp            string
	EmergencyName   string
	EmergencyStatus string
	EmergencyPhone  string
}

type UserDetailEntity struct {
	ID              uint
	UserID          uint   `validate:"required"`
	Fullame         string `validate:"required"`
	Email           string `validate:"required"`
	Password        string
	Address         string
	Gender          string
	UrlPhoto        string
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

type UserDataInterface interface {
	Insert(input Core, companyId int) error
	Login(email string, password string) (dataLogin Core, err error)
	SelectProfile(id int) (Core, error)
	UpdateProfile(id int, input UserDetailEntity) error
	UpdateOtherProfile(id int, input UserDetailEntity) error
	SelectAll(pageNumber int, pageSize int, ManagerId int, CompanyId int, filterManager int) ([]Core, error)
	SelectOtherProfile(id int) (Core, error)
	DeleteOtherProfile(id uint) error
}

type UserServiceInterface interface {
	Create(input Core, companyId int) error
	Login(email string, password string) (dataLogin Core, token string, err error)
	GetProfile(id int) (Core, error)
	UpdateProfile(id int, input UserDetailEntity) error
	UpdateOtherProfile(id int, input UserDetailEntity) error
	GetAll(pageNumber int, pageSize int, ManagerId int, CompanyId int, filterManager int) ([]Core, error)
	SelectOtherProfile(id int) (Core, error)
	DeleteOtherProfile(id uint) error
}

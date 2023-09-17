package employeeLevel

type Core struct {
	ID    uint
	Level string
}

type EmployeeLevelDataInterface interface {
	Insert(input Core) error
	SelectAll() ([]Core, error)
	EditEmployeeLevel(idEmployeeLevel uint, input Core) error
	DeleteEmployeeLevel(idEmployeeLevel int) error
}

type EmployeeLevelServiceInterface interface {
	Create(input Core) error
	GetAll() ([]Core, error)
	UpdateEmployeeLevel(idEmployeeLevel uint, input Core) error
	DeleteEmployeeLevel(idEmployeeLevel int) error
}

package role

type Core struct {
	ID       uint
	RoleName string
}

type RoleDataInterface interface {
	Insert(input Core) error
	SelectAll() ([]Core, error)
	EditRole(idRole uint, input Core) error
	DeleteRole(idRole int) error
}

type RoleServiceInterface interface {
	Create(input Core) error
	GetAll() ([]Core, error)
	UpdateRole(idRole uint, input Core) error
	DeleteRole(idRole int) error
}

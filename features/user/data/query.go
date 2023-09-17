package data

import (
	"errors"
	"fmt"
	"group-project-3/app/middlewares"
	_company "group-project-3/features/company/data"
	_level "group-project-3/features/employeeLevel/data"
	_role "group-project-3/features/role/data"
	"group-project-3/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(input user.Core) error {
	// mapping dari struct core to struct gorm model
	hashedPassword, _ := middlewares.HashedPassword(input.Password)
	input.Password = hashedPassword
	if input.UrlPhoto == "" {
		input.UrlPhoto = "https://ui-avatars.com/api/?name=" + input.Fullame
	}

	userGorm := CoreToModel(input)

	// simpan ke DB
	tx := repo.db.Create(&userGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (dataLogin user.Core, err error) {
	var data User
	var dataRole _role.Role
	var dataCompany _company.Company
	var dataLevel _level.EmployeeLevel

	repo.db.Raw("SELECT * FROM users WHERE email=?", email).Scan(&data)
	samePassword := middlewares.CheckPassword(password, data.Password)
	// hashedPassword, err := middlewares.HashedPassword(password)
	// fmt.Println("HASHED PASSWORD", hashedPassword)

	if samePassword {
		// tx := repo.db.Preload("Role").Where("email = ? AND password=?", email, password).Find(&data)
		query := `
		SELECT users.id,users.email,users.company_id,users.role_id,users.level_id,users.company_id ,roles.role_name,companies.name FROM users
		 INNER JOIN roles ON users.role_id= roles.id 
		 INNER JOIN companies ON users.company_id=companies.id 
		 WHERE users.email=? AND password=?
		`
		tx := repo.db.Raw(query, email, data.Password).Scan(&data)

		repo.db.Raw("SELECT *FROM roles WHERE id=?", data.RoleID).Scan(&dataRole)
		repo.db.Raw("SELECT *FROM companies WHERE id=?", data.CompanyID).Scan(&dataCompany)
		repo.db.Raw("SELECT *FROM employee_levels WHERE id=?", data.LevelID).Scan(&dataLevel)

		fmt.Println("LEVEL NAME ", dataLevel.Level)
		if tx.Error != nil {
			return user.Core{}, tx.Error
		}
		// if tx.Error != nil {
		// 	return
		// }
		// fmt.Println("DATA USERS DETAIL", dataUserDetail)

		if tx.RowsAffected == 0 {
			return user.Core{}, errors.New("data not found")
		}
		fmt.Println("LEVEL", dataLevel.Level)
		fmt.Println("Company", dataCompany.Name)
		fmt.Println("DARSADASTA", data)
		dataLogin = ModelToCore(data)
		dataLogin.Role.RoleName = dataRole.RoleName
		dataLogin.Company.CompanyName = dataCompany.Name
		dataLogin.Level.Level = dataLevel.Level
		fmt.Println("COA", dataLogin.Level.Level)
	} else {
		return user.Core{}, errors.New("data not found")
	}

	return dataLogin, nil
}

// SelectProfile implements user.UserDataInterface.
func (repo *userQuery) SelectProfile(id int) (user.Core, error) {
	var usersData User
	var roleData _role.Role
	tx := repo.db.First(&usersData, id).Scan(&usersData) // select * from users;
	txRole := repo.db.Raw("SELECT * FROM roles WHERE id=?", usersData.RoleID).Scan(&roleData)
	fmt.Println("ROLE DATA", roleData)

	fmt.Println("role", txRole)

	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user.Core{}, errors.New("data not found")
	}

	// if txRole.Error != nil {
	// 	return user.Core{}, tx.Error
	// }
	// if txRole.RowsAffected == 0 {
	// 	return user.Core{}, errors.New("data role not found")
	// }

	// fmt.Println("users:", usersData)
	//mapping dari struct gorm model ke struct core (entity)
	var usersCore = ModelToCore(usersData)

	return usersCore, nil
}

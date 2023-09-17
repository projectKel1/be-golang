package data

import (
	"errors"
	"fmt"
	"group-project-3/app/middlewares"
	_company "group-project-3/features/company/data"
	_level "group-project-3/features/employeeLevel/data"
	_role "group-project-3/features/role/data"
	_userDetail "group-project-3/features/userDetail/data"

	"group-project-3/features/user"

	// "group-project-3/features/user"
	// _userDetail "group-project-3/features/userDetail/data"

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
func (repo *userQuery) SelectProfile(id int) (dataProfile user.Core, err error) {
	var usersData User
	var roleData _role.Role
	var companyData _company.Company
	var levelData _level.EmployeeLevel
	var userDetailData _userDetail.UserDetail
	tx := repo.db.First(&usersData, id).Scan(&usersData) // select * from users;
	repo.db.Raw("SELECT * FROM roles WHERE id=?", usersData.RoleID).Scan(&roleData)
	repo.db.Raw("SELECT * FROM companies WHERE id=?", usersData.CompanyID).Scan(&companyData)
	repo.db.Raw("SELECT * FROM employee_levels WHERE id=?", usersData.CompanyID).Scan(&levelData)
	repo.db.Raw("SELECT * FROM user_details WHERE id=?", usersData.ID).Scan(&userDetailData)
	fmt.Println("user details", &userDetailData)
	// usersData.Role.RoleName = roleData.RoleName
	// dataProfile.Role.RoleName = roleData.RoleName
	// usersData.Role.RoleName = roleData.RoleName
	fmt.Println("roleku", usersData.Role)

	if tx.Error != nil {
		return dataProfile, tx.Error
	}
	if tx.RowsAffected == 0 {
		return dataProfile, errors.New("data not found")
	}

	//mapping dari struct gorm model ke struct core (entity)
	var usersCore = ModelToCore(usersData)
	usersCore.Role.RoleName = roleData.RoleName
	usersCore.Company.CompanyName = companyData.Name
	usersCore.Level.Level = levelData.Level
	usersCore.UserDetail.Nik = userDetailData.Nik
	usersCore.UserDetail.Address = userDetailData.Address
	usersCore.UserDetail.PhoneNumber = userDetailData.PhoneNumber
	usersCore.UserDetail.Gender = string(userDetailData.Gender)
	return usersCore, nil
}

// SelectAll implements user.UserDataInterface.
func (repo *userQuery) SelectAll(pageNumber int, pageSize int) ([]user.Core, error) {
	var userData []User
	var dataRole _role.Role
	var dataLevel _level.EmployeeLevel
	var dataCompany _company.Company
	offset := (pageNumber - 1) * pageSize

	tx := repo.db.Offset(offset).Limit(pageSize).Find(&userData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var userCore []user.Core

	for _, value := range userData {

		repo.db.Raw("SELECT *FROM roles WHERE id=?", value.RoleID).Scan(&dataRole)
		repo.db.Raw("SELECT *FROM employee_levels WHERE id=?", value.LevelID).Scan(&dataLevel)
		repo.db.Raw("SELECT *FROM companies WHERE id=?", value.CompanyID).Scan(&dataCompany)

		fmt.Println("LEVEL ", value.CreatedAt)
		userCore = append(userCore, user.Core{
			ID:          value.ID,
			Fullame:     value.Fullname,
			UrlPhoto:    value.UrlPhoto,
			RoleName:    dataRole.RoleName,
			LevelName:   dataLevel.Level,
			CompanyName: dataCompany.Name,
			Email:       value.Email,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}
	return userCore, nil

}

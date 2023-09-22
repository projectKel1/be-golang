package data

import (
	"errors"
	"fmt"
	"group-project-3/app/middlewares"
	_company "group-project-3/features/company/data"
	_level "group-project-3/features/employeeLevel/data"
	_role "group-project-3/features/role/data"
	"group-project-3/features/user"
	_userDetail "group-project-3/features/userDetail/data"

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
func (repo *userQuery) Insert(input user.Core, companyId int) error {
	// mapping dari struct core to struct gorm model
	hashedPassword, _ := middlewares.HashedPassword(input.Password)
	input.Password = hashedPassword
	input.CompanyID = uint(companyId)
	if input.UrlPhoto == "" {
		input.UrlPhoto = "https://ui-avatars.com/api/?name=" + input.Fullame
	}

	userGorm := CoreToModel(input)
	fmt.Println("cOMPANY ID,level id , manageer id", userGorm.CompanyID, userGorm.LevelID, userGorm.ManagerID)

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

	if samePassword {

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

		if tx.RowsAffected == 0 {
			return user.Core{}, errors.New("data not found")
		}

		fmt.Println("Company ID ", dataCompany.ID)
		dataLogin.CompanyID = dataCompany.ID
		dataLogin = ModelToCore(data)
		dataLogin.Role.RoleName = dataRole.RoleName
		dataLogin.Company.CompanyName = dataCompany.Name
		dataLogin.Company.ID = dataCompany.ID
		dataLogin.Level.Level = dataLevel.Level

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
	repo.db.Raw("SELECT * FROM user_details WHERE user_id=?", usersData.ID).Scan(&userDetailData)
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
	usersCore.UserDetail.NoKK = userDetailData.NoKK
	usersCore.UserDetail.Npwp = userDetailData.Npwp
	usersCore.UserDetail.Address = userDetailData.Address
	usersCore.UserDetail.PhoneNumber = userDetailData.PhoneNumber
	usersCore.UserDetail.Gender = string(userDetailData.Gender)
	usersCore.UserDetail.Bpjs = userDetailData.NoBPJS
	usersCore.UserDetail.EmergencyName = userDetailData.EmergencyName
	usersCore.UserDetail.EmergencyStatus = userDetailData.EmergencyStatus
	usersCore.UserDetail.EmergencyPhone = userDetailData.EmergencyPhone
	return usersCore, nil
}

// SelectAll implements user.UserDataInterface.
func (repo *userQuery) SelectAll(pageNumber int, pageSize int, managerId int, companyId int, filterManager int) ([]user.Core, error) {
	var userData []User
	var dataRole _role.Role
	var dataLevel _level.EmployeeLevel
	var dataCompany _company.Company
	offset := (pageNumber - 1) * pageSize

	var tx = repo.db
	if managerId != 0 {
		fmt.Println("Manager ID", managerId)
		tx = repo.db.Offset(offset).Limit(pageSize).Where("manager_id=? AND company_id=?", managerId, companyId).Find(&userData)
	} else if filterManager == 0 {
		tx = repo.db.Offset(offset).Limit(pageSize).Where("company_id=?", companyId).Find(&userData)
	} else {
		tx = repo.db.Offset(offset).Limit(pageSize).Where("company_id=? AND level_id = 2", companyId).Find(&userData)
	}

	if tx.Error != nil {
		return nil, tx.Error
	}

	var userCore []user.Core

	for _, value := range userData {

		repo.db.Raw("SELECT *FROM roles WHERE id=?", value.RoleID).Scan(&dataRole)
		repo.db.Raw("SELECT *FROM employee_levels WHERE id=?", value.LevelID).Scan(&dataLevel)
		repo.db.Raw("SELECT *FROM companies WHERE id=?", value.CompanyID).Scan(&dataCompany)

		fmt.Println("Status ", value.Status)
		userCore = append(userCore, user.Core{
			ID:          value.ID,
			Fullame:     value.Fullname,
			UrlPhoto:    value.UrlPhoto,
			RoleName:    dataRole.RoleName,
			LevelName:   dataLevel.Level,
			CompanyName: dataCompany.Name,
			Status:      value.Status,
			Email:       value.Email,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		})
	}
	return userCore, nil

}

// Update implements user.UserDataInterface.
func (repo *userQuery) UpdateProfile(id int, input user.UserDetailEntity) error {
	if input.UrlPhoto == "" {
		input.UrlPhoto = "https://ui-avatars.com/api/?name=" + input.Fullame
	}
	fmt.Println("EMERGENCY NAME", input.Nik)
	hashedPassword, _ := middlewares.HashedPassword(input.Password)

	tx := repo.db.Exec("update users SET fullname=?,email=?,password=?,url_photo=? WHERE id=?", input.Fullame, input.Email, hashedPassword, input.UrlPhoto, &id)
	var exists bool

	row := repo.db.Raw("SELECT*FROM user_details WHERE user_id=?", id)
	if errUserDetail := row.Scan(&exists); errUserDetail != nil {
		txDetail := repo.db.Exec("UPDATE user_details SET address=?,gender=?,phone_number=?,nik=?,no_kk=?,no_bpjs=?,npwp=?,emergency_name=?,emergency_status=?,emergency_phone=? WHERE user_id=?", input.Address, input.Gender, input.PhoneNumber, input.Nik, input.NoKK, input.NoBPJS, input.Npwp, input.EmergencyName, input.EmergencyStatus, input.EmergencyPhone, &id)
		if txDetail.Error != nil {
			return txDetail.Error
		}
	} else {
		txDetail := repo.db.Exec("INSERT INTO  user_details (user_id,address,gender,phone_number,nik,no_kk,no_bpjs,npwp,emergency_name,emergency_status,emergency_phone) VALUES(?,?,?,?,?,?,?,?,?,?,?)", &id, input.Address, input.Gender, input.PhoneNumber, input.Nik, input.NoKK, input.NoBPJS, input.Npwp, input.EmergencyName, input.EmergencyStatus, input.EmergencyPhone)

		if txDetail.Error != nil {
			return txDetail.Error
		}

	}

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Update implements user.UserDataInterface.
func (repo *userQuery) UpdateOtherProfile(id int, input user.UserDetailEntity) error {
	if input.UrlPhoto == "" {
		input.UrlPhoto = "https://ui-avatars.com/api/?name=" + input.Fullame
	}
	fmt.Println("EMERGENCY NAME", input.Nik)
	hashedPassword, _ := middlewares.HashedPassword(input.Password)
	tx := repo.db.Exec("update users SET fullname=?,email=?,password=?,url_photo=? WHERE id=?", input.Fullame, input.Email, hashedPassword, input.UrlPhoto, &id)

	var exists bool

	row := repo.db.Raw("SELECT*FROM user_details WHERE user_id=?", id)
	if errUserDetail := row.Scan(&exists); errUserDetail != nil {
		txDetail := repo.db.Exec("UPDATE user_details SET address=?,gender=?,phone_number=?,nik=?,no_kk=?,no_bpjs=?,npwp=?,emergency_name=?,emergency_status=?,emergency_phone=? WHERE user_id=?", input.Address, input.Gender, input.PhoneNumber, input.Nik, input.NoKK, input.NoBPJS, input.Npwp, input.EmergencyName, input.EmergencyStatus, input.EmergencyPhone, &id)

		if txDetail.Error != nil {
			return txDetail.Error
		}
	} else {
		txDetail := repo.db.Exec("INSERT INTO  user_details (user_id,address,gender,phone_number,nik,no_kk,no_bpjs,npwp,emergency_name,emergency_status,emergency_phone) VALUES(?,?,?,?,?,?,?,?,?,?,?)", &id, input.Address, input.Gender, input.PhoneNumber, input.Nik, input.NoKK, input.NoBPJS, input.Npwp, input.EmergencyName, input.EmergencyStatus, input.EmergencyPhone)

		if txDetail.Error != nil {
			return txDetail.Error
		}

	}

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteOtherProfile implements user.UserDataInterface.
func (repo *userQuery) DeleteOtherProfile(id uint) error {
	var userGorm User
	var userDetailGorm _userDetail.UserDetail
	tx := repo.db.Where("id = ?", id).Delete(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}

	var exists bool

	row := repo.db.Raw("SELECT*FROM user_details WHERE user_id=?", id)

	if errUserDetail := row.Scan(&exists); errUserDetail != nil {
		repo.db.Where("user_id = ?", id).Delete(&userDetailGorm)
	} else {
		fmt.Println("No data user details")
	}
	return nil
}

// SelectById implements user.UserDataInterface.
func (repo *userQuery) SelectById(id int) (userDetail user.UserDetailEntity, err error) {
	var userData User
	tx := repo.db.Raw("SELECT *FROM users WHERE id=?", id).Scan(&userData)
	if tx.Error != nil {
		return user.UserDetailEntity{}, tx.Error
	}
	fmt.Println("USER DATA", userData)
	if tx.RowsAffected == 0 {
		return user.UserDetailEntity{}, errors.New("data not found")
	}
	userDetail.ID = userData.ID
	userDetail.Email = userData.Email

	fmt.Println("user detail", userDetail)

	return userDetail, nil
}

// SelectOtherProfile implements user.UserDataInterface.
func (repo *userQuery) SelectOtherProfile(id int) (dataProfile user.Core, err error) {
	var usersData User
	var roleData _role.Role
	var companyData _company.Company
	var levelData _level.EmployeeLevel
	var userDetailData _userDetail.UserDetail
	tx := repo.db.Raw("SELECT*FROM users WHERE id=?", id).Scan(&usersData) // select * from users;

	repo.db.Raw("SELECT * FROM roles WHERE id=?", usersData.RoleID).Scan(&roleData)
	repo.db.Raw("SELECT * FROM companies WHERE id=?", usersData.CompanyID).Scan(&companyData)
	repo.db.Raw("SELECT * FROM employee_levels WHERE id=?", usersData.CompanyID).Scan(&levelData)
	repo.db.Raw("SELECT * FROM user_details WHERE user_id=?", usersData.ID).Scan(&userDetailData)
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
	usersCore.Company.ID = companyData.ID
	usersCore.Company.CompanyName = companyData.Name
	usersCore.Level.Level = levelData.Level
	usersCore.UserDetail.Nik = userDetailData.Nik
	usersCore.UserDetail.NoKK = userDetailData.NoKK
	usersCore.UserDetail.Npwp = userDetailData.Npwp
	usersCore.UserDetail.Address = userDetailData.Address
	usersCore.UserDetail.PhoneNumber = userDetailData.PhoneNumber
	usersCore.UserDetail.Gender = string(userDetailData.Gender)
	usersCore.UserDetail.Bpjs = userDetailData.NoBPJS
	usersCore.UserDetail.EmergencyName = userDetailData.EmergencyName
	usersCore.UserDetail.EmergencyStatus = userDetailData.EmergencyStatus
	usersCore.UserDetail.EmergencyPhone = userDetailData.EmergencyPhone
	return usersCore, nil
}

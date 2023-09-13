package data

import (
	"errors"
	"fmt"
	"group-project-3/app/middlewares"
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
	repo.db.Raw("SELECT id,role_id, email,password FROM users WHERE email=?", email).Scan(&data)
	samePassword := middlewares.CheckPassword(password, data.Password)
	// hashedPassword, err := middlewares.HashedPassword(password)
	// fmt.Println("HASHED PASSWORD", hashedPassword)

	if samePassword {
		tx := repo.db.Raw("SELECT id,role_id,status,company_id,url_photo,gender, email,password FROM users WHERE email=? AND password=?", email, data.Password).Scan(&data)
		if tx.Error != nil {
			return user.Core{}, tx.Error
		}
		fmt.Println("SAME PASSWORD", data.RoleID)
		if tx.RowsAffected == 0 {
			return user.Core{}, errors.New("data not found")
		}
		dataLogin = ModelToCore(data)
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

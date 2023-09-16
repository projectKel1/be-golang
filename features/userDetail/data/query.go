package data

import (
	"errors"
	"group-project-3/exception"
	"group-project-3/features/userDetail"
	"log"

	"gorm.io/gorm"
)

type UserDetailDataImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) userDetail.UserDetailDataInterface {
	return &UserDetailDataImpl{db: db}
}

// Delete implements userDetail.UserDetailDataInterface
func (data *UserDetailDataImpl) Delete(id uint) error {
	tx := data.db.Where("user_id = ?", id).Delete(&UserDetail{})

	if tx.Error != nil {
		return exception.ErrInternalServer
	}
	if tx.RowsAffected == 0 {
		return exception.ErrIdIsNotFound
	}

	return nil

}

// FindById implements userDetail.UserDetailDataInterface
func (data *UserDetailDataImpl) FindById(id uint) (userDetail.UserDetailEntity, error) {
	var userGorm UserDetail
	tx := data.db.Where("user_id = ?", id).First(&userGorm)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return userDetail.UserDetailEntity{}, exception.ErrIdIsNotFound
		} else {
			return userDetail.UserDetailEntity{}, exception.ErrInternalServer
		}
	}
	userEntity := ModelToCore(userGorm)
	return userEntity, nil
}

// Save implements userDetail.UserDetailDataInterface
func (data *UserDetailDataImpl) Save(userDetail userDetail.UserDetailEntity) error {
	userDetailGorm := CoreToModel(userDetail)
	tx := data.db.Create(&userDetailGorm)
	if tx.Error != nil {
		log.Println(tx.Error)
		return exception.ErrInternalServer
	}
	return nil
}

// Update implements userDetail.UserDetailDataInterface
func (data *UserDetailDataImpl) Update(userDetail userDetail.UserDetailEntity) error {
	var userDetailGorm UserDetail

	tx := data.db.Where("user_id = ?", userDetail.UserID).First(&userDetailGorm)
	if tx.Error != nil {
		return exception.ErrInternalServer
	}

	if tx.RowsAffected == 0 {
		return exception.ErrIdIsNotFound
	}
	userUpdate := CoreToModel(userDetail)
	tx = data.db.Model(&userDetailGorm).Updates(userUpdate)
	if tx.Error != nil {
		return exception.ErrInternalServer
	}

	if tx.RowsAffected == 0 {
		return exception.ErrIdIsNotFound
	}

	return nil
}

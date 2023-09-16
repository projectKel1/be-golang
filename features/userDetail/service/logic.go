package service

import "group-project-3/features/userDetail"

type UserDetailServiceImpl struct {
	UserDetailData userDetail.UserDetailDataInterface
}

func New(userDetailData userDetail.UserDetailDataInterface) userDetail.UserDetailServiceInterface {
	return &UserDetailServiceImpl{UserDetailData: userDetailData}
}

// Create implements userDetail.UserDetailServiceInterface
func (service *UserDetailServiceImpl) Create(userDetail userDetail.UserDetailEntity) error {
	err := service.UserDetailData.Save(userDetail)
	return err
}

// Delete implements userDetail.UserDetailServiceInterface
func (service *UserDetailServiceImpl) Delete(id uint) error {
	err := service.UserDetailData.Delete(id)
	return err
}

// FindById implements userDetail.UserDetailServiceInterface
func (service *UserDetailServiceImpl) FindById(id uint) (userDetail.UserDetailEntity, error) {
	result, err := service.UserDetailData.FindById(id)
	return result, err
}

// Update implements userDetail.UserDetailServiceInterface
func (service *UserDetailServiceImpl) Update(userDetail userDetail.UserDetailEntity) error {
	err := service.UserDetailData.Update(userDetail)
	return err
}

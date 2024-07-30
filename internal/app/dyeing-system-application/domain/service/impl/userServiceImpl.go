package impl

import (
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/domain/service"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"
)

type UserService struct {
	iUserRepo service.IUserRepository
}

var _ service.IUserService = &UserService{}

// AddUser implements service.IUserService.
func (u *UserService) AddUser(reqUser *entity.User) (int64, error) {
	return u.iUserRepo.AddUser(reqUser)
}

// DeleteUser implements service.IUserService.
func (u *UserService) DeleteUser(reqUser *entity.User) (int64, error) {
	return u.iUserRepo.DeleteUser(reqUser)
}

// QueryUser implements service.IUserService.
func (u *UserService) QueryUser(reqUser *entity.User) (*entity.User, error) {
	return u.iUserRepo.QueryUser(reqUser)
}

// UpdateUser implements service.IUserService.
func (u *UserService) UpdateUser(reqUser *entity.User) (int64, error) {
	return u.iUserRepo.UpdateUser(reqUser)
}

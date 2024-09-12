package repository

import (
	"errors"

	"gorm.io/gorm"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/domain/service"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"
)

type UserRepo struct {
	db *gorm.DB
}

// UserRepo implements the repository.UserRepository interface
var _ service.IUserRepository = &UserRepo{}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

// AddUser implements domain.IUserRepository.
func (r *UserRepo) AddUser(reqUser *entity.User) (int64, error) {
	result := r.db.Create(reqUser)
	return result.RowsAffected, result.Error
}

// DeleteUser implements domain.IUserRepository.
func (*UserRepo) DeleteUser(reqUser *entity.User) (int64, error) {
	panic("unimplemented")
}

// QueryUser implements domain.IUserRepository.
func (r *UserRepo) QueryUser(reqUser *entity.User) (*entity.User, error) {
	var user entity.User
	// err := r.db.Debug().Take(&user).Error
	result := r.db.Debug().Where("name = ? and password = ?", reqUser.Name, reqUser.Password).Take(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &user, nil
}

// UpdateUser implements domain.IUserRepository.
func (r *UserRepo) UpdateUser(reqUser *entity.User) (int64, error) {
	result := r.db.Model(&reqUser).Where("Name = ?", reqUser.Name).Updates(map[string]interface{}{"Password": reqUser.Password, "Avatar": reqUser.Avatar, "address": reqUser.Address, "phone": reqUser.Phone, "email": reqUser.Email})
	return result.RowsAffected, result.Error
}

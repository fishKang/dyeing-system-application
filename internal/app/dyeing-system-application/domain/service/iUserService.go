package service

import "wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"

type IUserService interface {
	QueryUser(reqUser *entity.User) (*entity.User, error)
	UpdateUser(reqUser *entity.User) (int64, error)
	DeleteUser(reqUser *entity.User) (int64, error)
	AddUser(reqUser *entity.User) (int64, error)
}

type IUserRepository interface {
	QueryUser(reqUser *entity.User) (*entity.User, error)
	UpdateUser(reqUser *entity.User) (int64, error)
	DeleteUser(reqUser *entity.User) (int64, error)
	AddUser(reqUser *entity.User) (int64, error)
}

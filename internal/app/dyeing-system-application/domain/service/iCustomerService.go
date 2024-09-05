package service

import "wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"

type ICustomerService interface {
	QueryCustomer(reqCustomer *entity.Customer) (*[]entity.Customer, error)
	UpdateCustomer(reqCustomer *entity.Customer) (int64, error)
	DeleteCustomer(reqCustomer *entity.Customer) (int64, error)
	AddCustomer(reqCustomer *entity.Customer) (int64, error)
}

type ICustomerRepository interface {
	QueryCustomer(reqCustomer *entity.Customer) (*[]entity.Customer, error)
	UpdateCustomer(reqCustomer *entity.Customer) (int64, error)
	DeleteCustomer(reqCustomer *entity.Customer) (int64, error)
	AddCustomer(reqCustomer *entity.Customer) (int64, error)
}

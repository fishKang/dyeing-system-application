package impl

import (
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/domain/service"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"
)

type CustomerService struct {
	iCustomerRepo service.ICustomerRepository
}

var _ service.ICustomerService = &CustomerService{}

// AddCustomer implements service.ICustomerService.
func (u *CustomerService) AddCustomer(reqCustomer *entity.Customer) (int64, error) {
	return u.iCustomerRepo.AddCustomer(reqCustomer)
}

// DeleteCustomer implements service.ICustomerService.
func (u *CustomerService) DeleteCustomer(reqCustomer *entity.Customer) (int64, error) {
	return u.iCustomerRepo.DeleteCustomer(reqCustomer)
}

// QueryCustomer implements service.ICustomerService.
func (u *CustomerService) QueryCustomer(reqCustomer *entity.Customer) (*[]entity.Customer, error) {
	return u.iCustomerRepo.QueryCustomer(reqCustomer)
}

// UpdateCustomer implements service.ICustomerService.
func (u *CustomerService) UpdateCustomer(reqCustomer *entity.Customer) (int64, error) {
	return u.iCustomerRepo.UpdateCustomer(reqCustomer)
}

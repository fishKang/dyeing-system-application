package repository

import (
	"errors"

	"gorm.io/gorm"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/domain/service"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"
)

type CustomerRepo struct {
	db *gorm.DB
}

// CustomerRepo implements the repository.CustomerRepository interface
var _ service.ICustomerRepository = &CustomerRepo{}

func NewCustomerRepo(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{db}
}

// AddCustomer implements domain.ICustomerRepository.
func (r *CustomerRepo) AddCustomer(reqCustomer *entity.Customer) (int64, error) {
	result := r.db.Create(reqCustomer)
	return result.RowsAffected, result.Error
}

// DeleteCustomer implements domain.ICustomerRepository.
func (*CustomerRepo) DeleteCustomer(reqCustomer *entity.Customer) (int64, error) {
	panic("unimplemented")
}

// QueryCustomer implements domain.ICustomerRepository.
func (r *CustomerRepo) QueryCustomer(reqCustomer *entity.Customer) (*[]entity.Customer, error) {
	var customer []entity.Customer
	// err := r.db.Debug().Take(&Customer).Error
	if reqCustomer.Name == "" && reqCustomer.PersonName == "" && reqCustomer.Address == "" && reqCustomer.Phone == "" && reqCustomer.Status == 0 {
		result := r.db.Find(&customer)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	} else {
		// result := r.db.Find(&Customer)
		status := 1
		if reqCustomer.Status == 0 {
			status = 1
		}
		result := r.db.Where("name LIKE ? and person_name like ? and address like ? and phone like ? and Status = ?", "%"+reqCustomer.Name+"%", "%"+reqCustomer.PersonName+"%", "%"+reqCustomer.Address+"%", "%"+reqCustomer.Phone+"%", status).Find(&customer)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &customer, nil
}

// UpdateCustomer implements domain.ICustomerRepository.
func (r *CustomerRepo) UpdateCustomer(reqCustomer *entity.Customer) (int64, error) {
	result := r.db.Model(&reqCustomer).Where("Name = ?", reqCustomer.Name).Updates(map[string]interface{}{"person_name": reqCustomer.PersonName, "address": reqCustomer.Address, "phone": reqCustomer.Phone})
	return result.RowsAffected, result.Error
}

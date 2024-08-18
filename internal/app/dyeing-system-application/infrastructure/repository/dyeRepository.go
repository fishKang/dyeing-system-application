package repository

import (
	"errors"

	"gorm.io/gorm"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/domain/service"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"
)

type DyeRepo struct {
	db *gorm.DB
}

// DyeRepo implements the repository.DyeRepository interface
var _ service.IDyeRepository = &DyeRepo{}

func NewDyeRepo(db *gorm.DB) *DyeRepo {
	return &DyeRepo{db}
}

// AddDye implements domain.IDyeRepository.
func (*DyeRepo) AddDye(reqDye *entity.Dye) (int64, error) {
	panic("unimplemented")
}

// DeleteDye implements domain.IDyeRepository.
func (*DyeRepo) DeleteDye(reqDye *entity.Dye) (int64, error) {
	panic("unimplemented")
}

// QueryDye implements domain.IDyeRepository.
func (r *DyeRepo) QueryDye(reqDye *entity.Dye) (*[]entity.Dye, error) {
	var dye []entity.Dye
	// err := r.db.Debug().Take(&Dye).Error
	result := r.db.Find(&dye)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &dye, nil
}

// UpdateDye implements domain.IDyeRepository.
func (*DyeRepo) UpdateDye(Dye *entity.Dye) (int64, error) {
	panic("unimplemented")
}

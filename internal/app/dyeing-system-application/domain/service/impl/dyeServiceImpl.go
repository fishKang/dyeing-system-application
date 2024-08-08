package impl

import (
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/domain/service"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"
)

type DyeService struct {
	iDyeRepo service.IDyeRepository
}

var _ service.IDyeService = &DyeService{}

// AddDye implements service.IDyeService.
func (u *DyeService) AddDye(reqDye *entity.Dye) (int64, error) {
	return u.iDyeRepo.AddDye(reqDye)
}

// DeleteDye implements service.IDyeService.
func (u *DyeService) DeleteDye(reqDye *entity.Dye) (int64, error) {
	return u.iDyeRepo.DeleteDye(reqDye)
}

// QueryDye implements service.IDyeService.
func (u *DyeService) QueryDye(reqDye *entity.Dye) (*[]entity.Dye, error) {
	return u.iDyeRepo.QueryDye(reqDye)
}

// UpdateDye implements service.IDyeService.
func (u *DyeService) UpdateDye(reqDye *entity.Dye) (int64, error) {
	return u.iDyeRepo.UpdateDye(reqDye)
}

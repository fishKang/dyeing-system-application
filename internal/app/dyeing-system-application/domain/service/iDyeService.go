package service

import "wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"

type IDyeService interface {
	QueryDye(reqDye *entity.Dye) (*[]entity.Dye, error)
	UpdateDye(reqDye *entity.Dye) (int64, error)
	DeleteDye(reqDye *entity.Dye) (int64, error)
	AddDye(reqDye *entity.Dye) (int64, error)
}

type IDyeRepository interface {
	QueryDye(reqDye *entity.Dye) (*[]entity.Dye, error)
	UpdateDye(reqDye *entity.Dye) (int64, error)
	DeleteDye(reqDye *entity.Dye) (int64, error)
	AddDye(reqDye *entity.Dye) (int64, error)
}

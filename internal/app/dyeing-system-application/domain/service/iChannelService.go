package service

import "wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"

type IChannelService interface {
	QueryChannel(reqChannel *entity.Channel) (*[]entity.Channel, error)
	UpdateChannel(reqChannel *entity.Channel) (int64, error)
	DeleteChannel(reqChannel *entity.Channel) (int64, error)
	AddChannel(reqChannel *entity.Channel) (int64, error)
}

type IChannelRepository interface {
	QueryChannel(reqChannel *entity.Channel) (*[]entity.Channel, error)
	UpdateChannel(reqChannel *entity.Channel) (int64, error)
	DeleteChannel(reqChannel *entity.Channel) (int64, error)
	AddChannel(reqChannel *entity.Channel) (int64, error)
}

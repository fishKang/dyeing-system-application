package impl

import (
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/domain/service"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"
)

type ChannelService struct {
	iChannelRepo service.IChannelRepository
}

var _ service.IChannelService = &ChannelService{}

// AddChannel implements service.IChannelService.
func (u *ChannelService) AddChannel(reqChannel *entity.Channel) (int64, error) {
	return u.iChannelRepo.AddChannel(reqChannel)
}

// DeleteChannel implements service.IChannelService.
func (u *ChannelService) DeleteChannel(reqChannel *entity.Channel) (int64, error) {
	return u.iChannelRepo.DeleteChannel(reqChannel)
}

// QueryChannel implements service.IChannelService.
func (u *ChannelService) QueryChannel(reqChannel *entity.Channel) (*[]entity.Channel, error) {
	return u.iChannelRepo.QueryChannel(reqChannel)
}

// UpdateChannel implements service.IChannelService.
func (u *ChannelService) UpdateChannel(reqChannel *entity.Channel) (int64, error) {
	return u.iChannelRepo.UpdateChannel(reqChannel)
}

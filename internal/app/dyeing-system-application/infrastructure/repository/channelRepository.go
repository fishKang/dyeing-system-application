package repository

import (
	"errors"

	"gorm.io/gorm"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/domain/service"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"
)

type ChannelRepo struct {
	db *gorm.DB
}

// ChannelRepo implements the repository.ChannelRepository interface
var _ service.IChannelRepository = &ChannelRepo{}

func NewChannelRepo(db *gorm.DB) *ChannelRepo {
	return &ChannelRepo{db}
}

// AddChannel implements domain.IChannelRepository.
func (r *ChannelRepo) AddChannel(reqChannel *entity.Channel) (int64, error) {
	result := r.db.Create(reqChannel)
	return result.RowsAffected, result.Error
}

// DeleteChannel implements domain.IChannelRepository.
func (*ChannelRepo) DeleteChannel(reqChannel *entity.Channel) (int64, error) {
	panic("unimplemented")
}

// QueryChannel implements domain.IChannelRepository.
func (r *ChannelRepo) QueryChannel(reqChannel *entity.Channel) (*[]entity.Channel, error) {
	var channel = []entity.Channel{}
	// result := r.db.Debug().Take(&reqChannel).Error
	result := r.db.Debug().Where("serialNum = ?", reqChannel.SerialNum).Find(&channel)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &channel, nil
}

// UpdateChannel implements domain.IChannelRepository.
func (r *ChannelRepo) UpdateChannel(reqChannel *entity.Channel) (int64, error) {
	result := r.db.Model(&reqChannel).Update("response", reqChannel.Response)
	return result.RowsAffected, result.Error
}

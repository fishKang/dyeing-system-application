package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/common"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/domain/model"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/domain/service"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/util"
)

type UserService struct {
	iUserRepo    service.IUserRepository
	iChannelRepo service.IChannelRepository
}

func NewUserService(iUserRepo service.IUserRepository, iChannelRepo service.IChannelRepository) *UserService {
	return &UserService{
		iUserRepo:    iUserRepo,
		iChannelRepo: iChannelRepo,
	}
}

func (s *UserService) UserLogin(c *gin.Context) {
	//初始化日志对象
	log := util.NewSugarLogZap()
	//初始化入参
	userDTO := model.GetUserDTO(c, log)
	//生成唯一ID
	userDTO.Channel.ID = uint64(uuid.New().ID())
	//将入参转为JSON
	tmperr := util.RequestSugarPrintInfo(log, &userDTO.Channel, userDTO)
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, &userDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	//登记请求信息到channel
	count, err := s.iChannelRepo.AddChannel(&userDTO.Channel)
	if err != nil || count != 1 {
		util.ResponseSugarPrintInfo(log, &userDTO.Channel, common.CreateFailResponse(util.InsertFailed, "插入请求到channel失败", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.InsertFailed, "插入请求到channel失败", err.Error()))
		return
	}
	//用户登录
	user, err := s.iUserRepo.QueryUser(&userDTO.User)
	if err != nil {
		util.ResponseSugarPrintInfo(log, &userDTO.Channel, common.CreateFailResponse(util.RecordNotFound, "用户登录异常", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, "用户登录异常", err.Error()))
		return
	}
	//将出参转为JSON
	tmperr = util.ResponseSugarPrintInfo(log, &userDTO.Channel, common.CreateSuccessResponse(user))
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, &userDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	//更新返回信息到channel
	returnCount, err := s.iChannelRepo.UpdateChannel(&userDTO.Channel)
	if err != nil || returnCount != 1 {
		util.ResponseSugarPrintInfo(log, &userDTO.Channel, common.CreateFailResponse(util.UpdateFailed, "更新返回到channel失败", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.UpdateFailed, "更新返回到channel失败", err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(user))
}

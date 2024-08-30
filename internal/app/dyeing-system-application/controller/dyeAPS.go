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

type DyeService struct {
	iDyeRepo     service.IDyeRepository
	iChannelRepo service.IChannelRepository
}

func NewDyeService(iDyeRepo service.IDyeRepository, iChannelRepo service.IChannelRepository) *DyeService {
	return &DyeService{
		iDyeRepo:     iDyeRepo,
		iChannelRepo: iChannelRepo,
	}
}

func (s *DyeService) QueryDyeList(c *gin.Context) {
	//初始化日志对象
	log := util.NewSugarLogZap()
	//初始化入参
	dyeDTO := model.GetDyeDTO(c, log)
	//生成唯一ID
	dyeDTO.Channel.SerialNum = uuid.New().String()
	//将入参转为JSON
	tmperr := util.RequestSugarPrintInfo(log, &dyeDTO.Channel, dyeDTO)
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, &dyeDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	//登记请求信息到channel
	count, err := s.iChannelRepo.AddChannel(&dyeDTO.Channel)
	if err != nil || count != 1 {
		util.ResponseSugarPrintInfo(log, &dyeDTO.Channel, common.CreateFailResponse(util.InsertFailed, "插入请求到channel失败", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.InsertFailed, "插入请求到channel失败", err.Error()))
		return
	}
	//查询染料清单
	dye, err := s.iDyeRepo.QueryDye(&dyeDTO.Dye)
	if err != nil {
		util.ResponseSugarPrintInfo(log, &dyeDTO.Channel, common.CreateFailResponse(util.RecordNotFound, "查询染料清单异常", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, "查询染料清单异常", err.Error()))
		return
	}

	// outputDye := []entity.Dye{}
	//将出参转为JSON
	tmperr = util.ResponseSugarPrintInfo(log, &dyeDTO.Channel, common.CreateSuccessResponse(dye))
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, &dyeDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	//更新返回信息到channel
	returnCount, err := s.iChannelRepo.UpdateChannel(&dyeDTO.Channel)
	if err != nil || returnCount != 1 {
		util.ResponseSugarPrintInfo(log, &dyeDTO.Channel, common.CreateFailResponse(util.UpdateFailed, "更新返回到channel失败", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.UpdateFailed, "更新返回到channel失败", err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(dye))
}

func (s *DyeService) UpdateDyeDetail(c *gin.Context) {
	//初始化日志对象
	log := util.NewSugarLogZap()
	//初始化入参
	dyeDTO := model.GetDyeDTO(c, log)
	//生成唯一ID
	dyeDTO.Channel.SerialNum = uuid.New().String()
	//将入参转为JSON
	tmperr := util.RequestSugarPrintInfo(log, &dyeDTO.Channel, dyeDTO)
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, &dyeDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	//更新染料明细
	count, err := s.iDyeRepo.UpdateDye(&dyeDTO.Dye)
	if err != nil || count != 1 {
		util.ResponseSugarPrintInfo(log, &dyeDTO.Channel, common.CreateFailResponse(util.RecordNotFound, "更新染料明细异常", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, "更新染料明细异常", err.Error()))
		return
	}

	// outputDye := []entity.Dye{}
	//将出参转为JSON
	tmperr = util.ResponseSugarPrintInfo(log, &dyeDTO.Channel, common.CreateSuccessResponse(dyeDTO.Dye))
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, &dyeDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(dyeDTO.Dye))
}

func (s *DyeService) AddDyeDetail(c *gin.Context) {
	//初始化日志对象
	log := util.NewSugarLogZap()
	//初始化入参
	dyeDTO := model.GetDyeDTO(c, log)
	//生成唯一ID
	dyeDTO.Channel.SerialNum = uuid.New().String()
	//将入参转为JSON
	tmperr := util.RequestSugarPrintInfo(log, &dyeDTO.Channel, dyeDTO)
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, &dyeDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	//更新染料明细
	count, err := s.iDyeRepo.UpdateDye(&dyeDTO.Dye)
	if err != nil || count != 1 {
		util.ResponseSugarPrintInfo(log, &dyeDTO.Channel, common.CreateFailResponse(util.RecordNotFound, "更新染料明细异常", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, "更新染料明细异常", err.Error()))
		return
	}

	// outputDye := []entity.Dye{}
	//将出参转为JSON
	tmperr = util.ResponseSugarPrintInfo(log, &dyeDTO.Channel, common.CreateSuccessResponse(dyeDTO.Dye))
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, &dyeDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(dyeDTO.Dye))
}

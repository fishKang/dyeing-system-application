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
	iDyeRepo service.IDyeService
}

func NewDyeService(iDyeRepo service.IDyeService) *DyeService {
	return &DyeService{
		iDyeRepo: iDyeRepo,
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
	tmperr := util.RequestSugarPrintInfo(log, dyeDTO.Channel, dyeDTO)
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, dyeDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	//查询染料清单
	dye, err := s.iDyeRepo.QueryDye(&dyeDTO.Dye)
	if err != nil {
		util.ResponseSugarPrintInfo(log, dyeDTO.Channel, common.CreateFailResponse(util.RecordNotFound, "查询染料清单异常", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, "查询染料清单异常", err.Error()))
		return
	}

	// outputDye := []entity.Dye{}
	//将出参转为JSON
	tmperr = util.ResponseSugarPrintInfo(log, dyeDTO.Channel, common.CreateSuccessResponse(dye))
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, dyeDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(dye))
}

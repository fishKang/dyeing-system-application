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

type CustomerService struct {
	iCustomerRepo service.ICustomerRepository
	iChannelRepo  service.IChannelRepository
}

func NewCustomerService(iCustomerRepo service.ICustomerRepository, iChannelRepo service.IChannelRepository) *CustomerService {
	return &CustomerService{
		iCustomerRepo: iCustomerRepo,
		iChannelRepo:  iChannelRepo,
	}
}

func (s *CustomerService) QueryCustomerList(c *gin.Context) {
	//初始化日志对象
	log := util.NewSugarLogZap()
	//初始化入参
	customerDTO := model.GetCustomerDTO(c, log)
	//生成唯一ID
	customerDTO.Channel.SerialNum = uuid.New().String()
	//将入参转为JSON
	tmperr := util.RequestSugarPrintInfo(log, &customerDTO.Channel, customerDTO)
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	//登记请求信息到channel
	count, err := s.iChannelRepo.AddChannel(&customerDTO.Channel)
	if err != nil || count != 1 {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.InsertFailed, "插入请求到channel失败", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.InsertFailed, "插入请求到channel失败", err.Error()))
		return
	}
	//查询客户信息清单
	customer, err := s.iCustomerRepo.QueryCustomer(&customerDTO.Customer)
	if err != nil {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.RecordNotFound, "查询客户信息清单异常", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, "查询客户信息清单异常", err.Error()))
		return
	}

	// outputCustomer := []entity.Customer{}
	//将出参转为JSON
	tmperr = util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateSuccessResponse(customer))
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	//更新返回信息到channel
	returnCount, err := s.iChannelRepo.UpdateChannel(&customerDTO.Channel)
	if err != nil || returnCount != 1 {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.UpdateFailed, "更新返回到channel失败", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.UpdateFailed, "更新返回到channel失败", err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(customer))
}

func (s *CustomerService) UpdateCustomerDetail(c *gin.Context) {
	//初始化日志对象
	log := util.NewSugarLogZap()
	//初始化入参
	customerDTO := model.GetCustomerDTO(c, log)
	//生成唯一ID
	customerDTO.Channel.SerialNum = uuid.New().String()
	//将入参转为JSON
	tmperr := util.RequestSugarPrintInfo(log, &customerDTO.Channel, customerDTO)
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	//登记请求信息到channel
	channelCount, err := s.iChannelRepo.AddChannel(&customerDTO.Channel)
	if err != nil || channelCount != 1 {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.InsertFailed, "插入请求到channel失败", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.InsertFailed, "插入请求到channel失败", err.Error()))
		return
	}
	//更新客户信息明细
	count, err := s.iCustomerRepo.UpdateCustomer(&customerDTO.Customer)
	if err != nil || count != 1 {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.RecordNotFound, "更新客户信息明细异常", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, "更新客户信息明细异常", err.Error()))
		return
	}

	// outputCustomer := []entity.Customer{}
	//将出参转为JSON
	tmperr = util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateSuccessResponse(customerDTO.Customer))
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	//更新返回信息到channel
	returnCount, err := s.iChannelRepo.UpdateChannel(&customerDTO.Channel)
	if err != nil || returnCount != 1 {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.UpdateFailed, "更新返回到channel失败", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.UpdateFailed, "更新返回到channel失败", err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(customerDTO.Customer))
}

func (s *CustomerService) AddCustomerDetail(c *gin.Context) {
	//初始化日志对象
	log := util.NewSugarLogZap()
	//初始化入参
	customerDTO := model.GetCustomerDTO(c, log)
	//生成唯一ID
	customerDTO.Channel.SerialNum = uuid.New().String()
	//将入参转为JSON
	tmperr := util.RequestSugarPrintInfo(log, &customerDTO.Channel, customerDTO)
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	//登记请求信息到channel
	channelCount, err := s.iChannelRepo.AddChannel(&customerDTO.Channel)
	if err != nil || channelCount != 1 {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.InsertFailed, "插入请求到channel失败", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.InsertFailed, "插入请求到channel失败", err.Error()))
		return
	}
	//新增客户信息明细
	count, err := s.iCustomerRepo.AddCustomer(&customerDTO.Customer)
	if err != nil || count != 1 {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.RecordNotFound, "新增客户信息明细异常", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, "新增客户信息明细异常", err.Error()))
		return
	}

	// outputCustomer := []entity.Customer{}
	//将出参转为JSON
	tmperr = util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateSuccessResponse(customerDTO.Customer))
	if tmperr != nil {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.RecordNotFound, tmperr.Error(), tmperr.Error()))
		return
	}
	//更新返回信息到channel
	returnCount, err := s.iChannelRepo.UpdateChannel(&customerDTO.Channel)
	if err != nil || returnCount != 1 {
		util.ResponseSugarPrintInfo(log, &customerDTO.Channel, common.CreateFailResponse(util.UpdateFailed, "更新返回到channel失败", err.Error()))
		c.JSON(http.StatusOK, common.CreateFailResponse(util.UpdateFailed, "更新返回到channel失败", err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(customerDTO.Customer))
}

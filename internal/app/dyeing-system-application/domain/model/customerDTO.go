package model

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"
)

type CustomerDTO struct {
	entity.Customer `json:"data"`
	entity.Channel  `json:"channel"`
}

func GetCustomerDTO(c *gin.Context, log *zap.SugaredLogger) CustomerDTO {
	var customerDTO = CustomerDTO{}
	if errA := c.ShouldBindBodyWith(&customerDTO, binding.JSON); errA != nil {
		c.JSON(http.StatusBadRequest, errA.Error())
	}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Panic("CustomerDTO解析异常")
	}
	customerDTO.Channel.Request = string(body)
	return customerDTO
}

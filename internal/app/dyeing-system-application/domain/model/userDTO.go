package model

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"
)

type UserDTO struct {
	entity.User    `json:"data"`
	entity.Channel `json:"channel"`
}

func GetUserDTO(c *gin.Context, log *zap.SugaredLogger) UserDTO {
	var userDTO = UserDTO{}
	if errA := c.ShouldBindBodyWith(&userDTO, binding.JSON); errA != nil {
		c.JSON(http.StatusBadRequest, errA.Error())
	}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Panic("UserDTO解析异常")
	}
	userDTO.Channel.Request = string(body)
	return userDTO
}

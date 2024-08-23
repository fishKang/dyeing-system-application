package model

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/infrastructure/entity"
)

type DyeDTO struct {
	entity.Dye     `json:"data"`
	entity.Channel `json:"channel"`
}

func GetDyeDTO(c *gin.Context, log *zap.SugaredLogger) DyeDTO {
	var dyeDTO = DyeDTO{}
	if errA := c.ShouldBindBodyWith(&dyeDTO, binding.JSON); errA != nil {
		c.JSON(http.StatusBadRequest, errA.Error())
	}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Panic("DyeDTO解析异常")
	}
	dyeDTO.Channel.Request = string(body)
	return dyeDTO
}

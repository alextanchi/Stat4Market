package controller

import (
	"Stat4Market/internal/models"
	"Stat4Market/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Controller interface {
	CreateEvent(ctx *gin.Context)
}
type EventController struct {
	useCase service.Service
}

func NewController(srv service.Service) Controller {
	return &EventController{
		useCase: srv,
	}
}

func (c EventController) CreateEvent(ctx *gin.Context) {
	event := models.Request{}
	err := ctx.ShouldBind(&event)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректные данные",
		})
		return
	}
	err = c.useCase.CreateEvent(ctx, event)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Внутренняя ошибка сервера",
		})
		return
	}
	ctx.JSON(http.StatusOK, "Событие создано")
	return
}

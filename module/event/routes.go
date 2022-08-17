package event

import (
	"scheduler/module/event/controller"

	"github.com/gin-gonic/gin"
)

func AvaiableRoutes(events *gin.RouterGroup) {
	events.GET("/", controller.Find)
	events.POST("/", controller.Create)
	events.PATCH("/", controller.Update)
	events.DELETE("/", controller.Delete)
}

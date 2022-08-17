package schedule

import (
	"scheduler/module/user/controller"

	"github.com/gin-gonic/gin"
)

func AvaiableRoutes(users *gin.RouterGroup) {
	users.GET("/", controller.Find)
	users.POST("/", controller.Create)
	users.PATCH("/", controller.Update)
	users.DELETE("/", controller.Delete)
}

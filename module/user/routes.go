package user

import (
	"github.com/gin-gonic/gin"

	"scheduler/module/user/controller"
)

func AvaiableRoutes(users *gin.RouterGroup) {
	users.GET("/Log-in", controller.Find)
	users.POST("/Sign-up", controller.Create)
	users.PATCH("/Profile", controller.Update)
	users.DELETE("/Profile", controller.Delete)
}

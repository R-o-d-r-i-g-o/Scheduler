package pastor

import (
	"github.com/gin-gonic/gin"

	"scheduler/module/pastor/controller"
)

func AvaiableRoutes(pastors *gin.RouterGroup) {
	pastors.GET("/List", controller.Find)
	pastors.POST("/Create", controller.Create)
	pastors.PATCH("/{}", controller.Update)
	pastors.DELETE("/{}", controller.Delete)
}

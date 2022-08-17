package router

import (
	"scheduler/module/event"
	"scheduler/module/pastor"
	"scheduler/module/schedule"
	"scheduler/module/user"

	"github.com/gin-gonic/gin"
)

func Avaible(r *gin.Engine) {

	webSite := r.Group("/Scheduler")
	{
		users := webSite.Group("/Users")
		user.AvaiableRoutes(users)

		events := webSite.Group("/Events")
		event.AvaiableRoutes(events)

		pastors := webSite.Group("/Pastors")
		{
			pastor.AvaiableRoutes(pastors)

			schedules := pastors.Group("/Schedule")
			schedule.AvaiableRoutes(schedules)
		}
	}
}

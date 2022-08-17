package controller

import (
	"scheduler/database/model"
	"scheduler/util"

	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	var registred model.UserReferences

	registred.Find()
}

func Create(ctx *gin.Context) {
	var new model.UserReferences

	if err := ctx.BindJSON(&new.User); err == nil && util.IsEmailValid(new.User.Email) {
		util.AcceptedOrNotStatusReturn(new.Create(), ctx, new.User)
	}
}

func Update(ctx *gin.Context) {

}

func Delete(ctx *gin.Context) {

}

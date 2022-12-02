package handlers

import (
	"github.com/TulioGuaraldoB/lemarchand-password/core/controllers"
	"github.com/gin-gonic/gin"
)

func NewUserHandler(ginRouter *gin.RouterGroup, userController controllers.IUserController) {
	user := ginRouter.Group("user")
	{
		user.POST("verify", userController.VerifyPassword)
	}
}

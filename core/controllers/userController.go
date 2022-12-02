package controllers

import (
	"net/http"

	"github.com/TulioGuaraldoB/lemarchand-password/core/businesses"
	"github.com/TulioGuaraldoB/lemarchand-password/core/dtos/requests"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	VerifyPassword(ctx *gin.Context)
}

type userController struct {
	userBusiness businesses.IUserBusiness
}

func NewUserController(userBusiness businesses.IUserBusiness) IUserController {
	return &userController{
		userBusiness: userBusiness,
	}
}

func (c *userController) VerifyPassword(ctx *gin.Context) {
	passwordReq := new(requests.PasswordRequest)
	if err := ctx.ShouldBindJSON(passwordReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.userBusiness.VerifyPassword(passwordReq)
}

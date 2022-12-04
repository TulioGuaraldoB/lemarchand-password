package routes

import (
	"github.com/TulioGuaraldoB/lemarchand-password/util"
	"github.com/gin-gonic/gin"
)

func SetRoutes() (*gin.Engine, *gin.RouterGroup) {
	router := gin.Default()
	apiVersion := new(gin.RouterGroup)

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		apiVersion = v1
	}

	router.GET("health", util.Health)

	return router, apiVersion
}

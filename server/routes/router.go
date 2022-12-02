package routes

import "github.com/gin-gonic/gin"

func SetRoutes() *gin.RouterGroup {
	router := gin.Default()
	apiVersion := new(gin.RouterGroup)

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		apiVersion = v1
	}

	return apiVersion
}

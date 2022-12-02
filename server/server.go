package server

import (
	"log"

	"github.com/TulioGuaraldoB/lemarchand-password/config/env"
	"github.com/TulioGuaraldoB/lemarchand-password/core/businesses"
	"github.com/TulioGuaraldoB/lemarchand-password/core/controllers"
	"github.com/TulioGuaraldoB/lemarchand-password/core/handlers"
	"github.com/TulioGuaraldoB/lemarchand-password/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	Server *gin.Engine
}

func New() Server {
	return Server{
		Port:   env.Env.Port,
		Server: gin.Default(),
	}
}

func (s *Server) Run() {
	// Services
	router := routes.SetRoutes()

	// Businesses
	userBusiness := businesses.NewUserBusiness()

	// Controllers
	userController := controllers.NewUserController(userBusiness)

	// Handlers
	handlers.NewUserHandler(router, userController)

	log.Fatal(s.Server.Run(":" + s.Port))
}

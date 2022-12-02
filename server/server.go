package server

import (
	"log"

	"github.com/TulioGuaraldoB/lemarchand-password/config/env"
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
	log.Fatal(s.Server.Run(":" + s.Port))
}

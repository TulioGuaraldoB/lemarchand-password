package main

import (
	"github.com/TulioGuaraldoB/lemarchand-password/config/env"
	"github.com/TulioGuaraldoB/lemarchand-password/server"
)

func main() {
	env.GetEnvironmentVariables()

	server := server.New()
	server.Run()
}

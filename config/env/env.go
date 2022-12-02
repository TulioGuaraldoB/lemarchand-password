package env

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVariable struct {
	Port string
}

var Env *EnvironmentVariable

func GetEnvironmentVariables() *EnvironmentVariable {
	godotenv.Load(".env")

	Env = &EnvironmentVariable{
		Port: os.Getenv("PORT"),
	}

	return Env
}

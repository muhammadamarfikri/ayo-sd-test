package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error when loading the env file")
	}
}

type App struct {
	Env                              string
	Port                             string
	Version                          string
	Address                          string
	BaseUrl                          string
	Timezone                         string
	EndpointPrefix                   string
	GracefulShutdownTimeoutInSeconds int
}

func AppServer() *App {
	return &App{
		Env:  os.Getenv("APP_ENV"),
		Port: os.Getenv("APP_PORT"),
	}
}

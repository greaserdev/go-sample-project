package main

import (
	"golang-oauth-tutor/config"
	"golang-oauth-tutor/internal/routes"
)

func main() {
	config.InitAppEnv()

	httpRouter := routes.HttpRouter()
	httpRouter.Run(":" + config.Env.AppPort)
}

package config

import (
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	AppPort            string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
}

var Env env

func InitAppEnv() {
	godotenv.Load()
	Env.AppPort = os.Getenv("APP_PORT")
	Env.GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	Env.GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	Env.GoogleRedirectURL = os.Getenv("GOOGLE_OAUTH_REDIRECT_URL")
}

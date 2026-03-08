package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOauthConfig struct {
	*oauth2.Config
}

func GoogleOauthConfigInit() *GoogleOauthConfig {
	googleOauthConfig := &oauth2.Config{
		RedirectURL:  Env.GoogleRedirectURL,
		ClientID:     Env.GoogleClientID,
		ClientSecret: Env.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	return &GoogleOauthConfig{googleOauthConfig}
}

package services

import (
	"context"
	"encoding/json"
	"golang-oauth-tutor/config"
	"golang-oauth-tutor/internal/domain"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

type OAuthService struct {
	GoogleAuthConfig *config.GoogleOauthConfig
}

func NewOAuthService(gac *config.GoogleOauthConfig) *OAuthService {
	return &OAuthService{GoogleAuthConfig: gac}
}

func (svc *OAuthService) ExchangeGoogleCode(context context.Context, code string) (*oauth2.Token, error) {
	token, err := svc.GoogleAuthConfig.Exchange(context, code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (svc *OAuthService) GetGoogleUserData(token string) (domain.GoogleUser, error) {
	var userInfo domain.GoogleUser
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if err != nil {
		return domain.GoogleUser{}, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &userInfo)
	return userInfo, nil
}

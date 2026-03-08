package controllers

import (
	"context"
	"fmt"
	"golang-oauth-tutor/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RedirectController struct {
	oAuthService *services.OAuthService
}

type LoginRequest struct {
	Code string `json:"code"`
}

func NewRedirectController(oAuthService *services.OAuthService) *RedirectController {
	return &RedirectController{oAuthService: oAuthService}
}
func (rc *RedirectController) Redirect(c *gin.Context) {
	token, err := rc.oAuthService.ExchangeGoogleCode(context.Background(), c.Query("code"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("Error" + err.Error())
		return
	}

	googleUser, err := rc.oAuthService.GetGoogleUserData(token.AccessToken)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": googleUser.Email})
}

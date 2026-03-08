//go:build wireinject
// +build wireinject

package wire

import (
	"golang-oauth-tutor/config"
	"golang-oauth-tutor/internal/controllers"
	"golang-oauth-tutor/internal/services"

	googleWire "github.com/google/wire"
)

var oauthSet = googleWire.NewSet(config.GoogleOauthConfigInit)
var serviceSet = googleWire.NewSet(services.NewOAuthService)

func InitRedirectController() *controllers.RedirectController {
	googleWire.Build(oauthSet, serviceSet, controllers.NewRedirectController)
	return &controllers.RedirectController{}
}

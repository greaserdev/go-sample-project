package routes

import (
	"golang-oauth-tutor/wire"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HttpRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group("/api/v1")
	{
		redirect := v1.Group("redirects")
		{
			redirectController := wire.InitRedirectController()
			redirect.GET("/oauth/google", redirectController.Redirect)
		}

	}

	return router
}

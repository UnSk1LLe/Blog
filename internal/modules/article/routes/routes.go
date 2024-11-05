package routes

import (
	"blogProject/internal/middleware"
	articleCrtl "blogProject/internal/modules/article/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	articleController := articleCrtl.New()
	router.GET("/articles/:id", articleController.Show)

	authGroup := router.Group("/articles")
	authGroup.Use(middleware.IsAuth())
	{
		authGroup.GET("/create", articleController.Create)
		authGroup.POST("/store", articleController.Store)
	}
}

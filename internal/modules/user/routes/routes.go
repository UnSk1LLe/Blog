package routes

import (
	"blogProject/internal/middleware"
	userCtrl "blogProject/internal/modules/user/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	userController := userCtrl.New()

	guestGroup := router.Group("/")
	guestGroup.Use(middleware.IsGuest())
	{
		guestGroup.GET("/register", userController.Register)
		guestGroup.POST("/register", userController.HandleRegister)

		guestGroup.GET("/login", userController.Login)
		guestGroup.POST("/login", userController.HandleLogin)
	}

	authGroup := router.Group("/")
	authGroup.Use(middleware.IsAuth())
	{
		authGroup.POST("/logout", userController.HandleLogout)
	}
}

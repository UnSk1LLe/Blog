package routes

import (
	homeCtrl "blogProject/internal/modules/controllers"
	"blogProject/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {
	homeController := homeCtrl.NewController()
	router.GET("/", homeController.Index)
	router.GET("/login", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/view/login", gin.H{
			"title": "Login Page",
		})
	})
	router.GET("/register", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/view/register", gin.H{
			"title": "Register Page",
		})
	})
}

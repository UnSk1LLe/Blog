package routes

import (
	homeRoutes "blogProject/internal/modules/article/routes"
	articleRoutes "blogProject/internal/modules/home/routes"
	userRoutes "blogProject/internal/modules/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.Routes(router)
}

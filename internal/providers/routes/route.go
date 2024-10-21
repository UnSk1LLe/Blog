package routes

import (
	homeRoutes "blogProject/internal/modules/home/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
}

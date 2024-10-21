package html

import (
	"blogProject/internal/providers/view"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, status int, name string, data gin.H) {
	data = view.WithGlobalData(data)
	c.HTML(status, name, data)
}

package html

import "github.com/gin-gonic/gin"

func LoadHtml(router *gin.Engine) {
	//internal/modules/moduleName/view/view.tmpl
	router.LoadHTMLGlob("internal/**/**/**/*tmpl")
}

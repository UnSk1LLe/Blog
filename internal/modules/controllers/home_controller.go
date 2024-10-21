package controllers

import (
	"blogProject/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/view/home", gin.H{
		"title": "Home Page",
	})
}

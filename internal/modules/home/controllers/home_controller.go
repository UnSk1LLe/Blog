package controllers

import (
	ArticleService "blogProject/internal/modules/article/services"
	"blogProject/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {

	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "Home Page",
		"featured": controller.articleService.GetFeaturedArticles(),
		"stories":  controller.articleService.GetStoriesArticles(),
	})
}

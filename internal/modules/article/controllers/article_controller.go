package controllers

import (
	"blogProject/internal/modules/article/request/articles"
	ArticleResponse "blogProject/internal/modules/article/responses"
	ArticleService "blogProject/internal/modules/article/services"
	"blogProject/internal/modules/user/helpers"
	"blogProject/pkg/converters"
	"blogProject/pkg/errors"
	"blogProject/pkg/html"
	"blogProject/pkg/old"
	"blogProject/pkg/sessions"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {
	//get the article id
	var article ArticleResponse.Article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{
			"title":   "Server Error",
			"message": "error converting the id",
		})
		return
	}
	//find the article from db
	article, err = controller.articleService.Find(id)

	//if not found show error page
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{
			"title":   "Page not found",
			"message": "Seems like this article do not exist.",
		})
		return
	}

	//if found render the template

	html.Render(c, http.StatusOK, "modules/articles/html/show", gin.H{
		"title":   "Article page",
		"article": article,
	})
}

func (controller *Controller) Create(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/articles/html/create", gin.H{
		"title": "Create Article",
	})
}

func (controller *Controller) Store(c *gin.Context) {
	//validate the request
	var request articles.StoreRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	user := helpers.Auth(c)

	article, err := controller.articleService.StoreAsUser(request, user)

	//check for errors during article creation
	if err != nil {
		c.Redirect(http.StatusFound, "/articles/create")
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	//after creating the user > redirect to home page
	log.Printf("The article with ID %d created successfully by User %s \n", article, user.Name)
	c.Redirect(http.StatusFound, fmt.Sprintf("/articles/%d", article.ID))
}

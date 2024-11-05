package services

import (
	"blogProject/internal/modules/article/request/articles"
	ArticleResponse "blogProject/internal/modules/article/responses"
	userResponse "blogProject/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
	StoreAsUser(request articles.StoreRequest, user userResponse.User) (ArticleResponse.Article, error)
}

package services

import (
	ArticleModel "blogProject/internal/modules/article/models"
	ArticleRepository "blogProject/internal/modules/article/repositories"
	"blogProject/internal/modules/article/request/articles"
	ArticleResponse "blogProject/internal/modules/article/responses"
	userResponse "blogProject/internal/modules/user/responses"
	"blogProject/pkg/database"
	"errors"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(database.Connection()),
	}
}

func (articleService *ArticleService) GetFeaturedArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(4)
	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleService) GetStoriesArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(6)
	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleService) Find(id int) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article
	article := articleService.articleRepository.Find(id)
	if article.ID == 0 {
		return response, errors.New("article not found")
	}
	response = ArticleResponse.ToArticle(article)
	return response, nil
}

func (articleService *ArticleService) StoreAsUser(request articles.StoreRequest, user userResponse.User) (ArticleResponse.Article, error) {
	var article ArticleModel.Article

	article.Title = request.Title
	article.Content = request.Content
	article.UserID = user.ID

	newArticle := articleService.articleRepository.Create(article)

	if newArticle.ID == 0 {
		return ArticleResponse.Article{}, errors.New("error creating article")
	}

	return ArticleResponse.ToArticle(newArticle), nil
}

package repositories

import ArticleModel "blogProject/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
	Find(id int) ArticleModel.Article
	Create(article ArticleModel.Article) ArticleModel.Article
}

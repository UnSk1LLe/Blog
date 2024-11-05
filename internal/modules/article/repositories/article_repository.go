package repositories

import (
	ArticleModel "blogProject/internal/modules/article/models"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{
		DB: db,
	}
}

func (a *ArticleRepository) List(limit int) []ArticleModel.Article {
	var articles []ArticleModel.Article
	a.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)
	return articles
}

func (a *ArticleRepository) Find(id int) ArticleModel.Article {
	var article ArticleModel.Article
	a.DB.Joins("User").First(&article, id)

	return article
}

func (a *ArticleRepository) Create(article ArticleModel.Article) ArticleModel.Article {
	var newArticle ArticleModel.Article

	a.DB.Create(&article).Scan(&newArticle)

	return newArticle
}

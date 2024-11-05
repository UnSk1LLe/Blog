package repositories

import (
	userModel "blogProject/internal/modules/user/models"
)

type ArticleRepositoryInterface interface {
	Create(user userModel.User) userModel.User
	FindByEmail(email string) userModel.User
	FindByID(id int) userModel.User
}

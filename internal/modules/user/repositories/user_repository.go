package repositories

import (
	userModel "blogProject/internal/modules/user/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) Create(user userModel.User) userModel.User {
	var newUser userModel.User
	u.DB.Create(&user).Scan(&newUser)
	return newUser
}

func (u *UserRepository) FindByEmail(email string) userModel.User {
	var user userModel.User
	u.DB.First(&user, "email = ?", email)
	return user
}

func (u *UserRepository) FindByID(id int) userModel.User {
	var user userModel.User
	u.DB.First(&user, "id = ?", id)
	return user
}

package models

import (
	"blogProject/internal/modules/user/models"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"varchar(255)"`
	Content string `gorm:"varchar(255)"`
	UserID  uint
	User    models.User
}

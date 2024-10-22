package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"varchar(255)"`
	Email    string `gorm:"varchar(255)"`
	Password string `gorm:"varchar(255)"`
}

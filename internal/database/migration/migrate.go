package migration

import (
	articleModels "blogProject/internal/modules/article/models"
	userModels "blogProject/internal/modules/user/models"
	"blogProject/pkg/database"
	"fmt"
	"log"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})

	if err != nil {
		log.Fatalf("Error migrating database: %v", err.Error())
	}

	fmt.Println("Database migration successfull")
}

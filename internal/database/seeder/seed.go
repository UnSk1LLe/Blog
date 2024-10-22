package seeder

import (
	ArticleModel "blogProject/internal/modules/article/models"
	userModel "blogProject/internal/modules/user/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

func Seed(db *gorm.DB) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("qwerty"), 12)
	if err != nil {
		log.Fatalf("error hashing password: %s", err.Error())
		return
	}
	user := userModel.User{Name: "Some name", Email: "example@gmail.com", Password: string(hashedPassword)}

	db.Create(&user)
	log.Printf("User with email: <%s> was CREATED!", user.Email)

	for i := 1; i <= 10; i++ {
		article := ArticleModel.Article{User: user, Title: fmt.Sprintf("Title %d", i), Content: fmt.Sprintf("Some random content %d", i), UserID: user.ID}
		db.Create(&article)

		log.Printf("Article with Title: <%s> was CREATED!", article.Title)
	}
	log.Printf("Seeder done ..")

}

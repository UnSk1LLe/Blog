package middleware

import (
	UserRepository "blogProject/internal/modules/user/repositories"
	"blogProject/pkg/database"
	"blogProject/pkg/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IsAuth() gin.HandlerFunc {
	var userRepo = UserRepository.New(database.Connection())

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)

		if user.ID == 0 {
			c.Redirect(http.StatusFound, "/login")
		}

		c.Next()
	}
}

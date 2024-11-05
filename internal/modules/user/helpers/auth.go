package helpers

import (
	UserRepository "blogProject/internal/modules/user/repositories"
	UserResponse "blogProject/internal/modules/user/responses"
	"blogProject/pkg/database"
	"blogProject/pkg/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Auth(c *gin.Context) UserResponse.User {
	var response UserResponse.User
	authID := sessions.Get(c, "auth")
	userID, err := strconv.Atoi(authID)
	if err != nil {
		return response
	}

	var userRepo = UserRepository.New(database.Connection())

	user := userRepo.FindByID(userID)
	if user.ID == 0 {
		return response
	}

	response = UserResponse.ToUser(user)

	return response
}

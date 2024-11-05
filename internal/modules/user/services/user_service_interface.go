package services

import (
	"blogProject/internal/modules/user/request/auth"
	userResponse "blogProject/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Create(request auth.RegisterRequest) (userResponse.User, error)
	CheckUserExists(email string) bool
	HandleUserLogin(request auth.LoginRequest) (userResponse.User, error)
}

package services

import (
	userModel "blogProject/internal/modules/user/models"
	UserRepository "blogProject/internal/modules/user/repositories"
	"blogProject/internal/modules/user/request/auth"
	UserResponse "blogProject/internal/modules/user/responses"
	"blogProject/pkg/database"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository UserRepository.ArticleRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: UserRepository.New(database.Connection()),
	}
}

func (u *UserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
	var response UserResponse.User
	var user userModel.User
	var err error

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return response, errors.New("error hashing password")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)

	newUser := u.userRepository.Create(user)
	if newUser.ID == 0 {
		return response, errors.New("error creating user")
	}

	return UserResponse.ToUser(newUser), nil
}

func (u *UserService) CheckUserExists(email string) bool {
	user := u.userRepository.FindByEmail(email)
	if user.ID != 0 {
		return true
	}
	return false
}

func (u *UserService) HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error) {
	var response UserResponse.User

	existsUser := u.userRepository.FindByEmail(request.Email)
	if existsUser.ID == 0 {
		return response, errors.New("could not find existsUser")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existsUser.Password), []byte(request.Password))
	if err != nil {
		return response, errors.New("wrong password")
	}

	response = UserResponse.ToUser(existsUser)

	return response, nil
}

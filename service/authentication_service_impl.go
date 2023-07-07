package service

import (
	"errors"
	"golang-jwttoken/config"
	"golang-jwttoken/data/request"
	"golang-jwttoken/helper"
	"golang-jwttoken/model"
	"golang-jwttoken/repository"
	"golang-jwttoken/utils"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewAuthenticationServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

// Login implements AuthenticationService.
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	// Find username in database
	new_user, user_err := a.UserRepository.FindByUserName(users.Username)
	if user_err != nil {
		return "", errors.New("invalid username or password")
	}

	config, _ := config.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_user.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_user.Id, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil
}

// Register implements AuthenticationService.
func (a *AuthenticationServiceImpl) Register(users request.CreateUsersRequest) (string, error) {
	_, new_user_error := a.UserRepository.FindByUserName(users.Username)
	if new_user_error == nil {
		return "", errors.New("username is existed")
	}
	hashedPassword, err := utils.HashPassword(users.Password)
	helper.ErrorPanic(err)

	newUser := model.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}
	a.UserRepository.Save(newUser)
	return "", nil
}

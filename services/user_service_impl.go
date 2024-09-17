package services

import (
	"errors"
	"github.com/jinzhu/copier"
	"log"
	"time"
	"todo-app/dto/request"
	"todo-app/dto/response"
	"todo-app/repositories"
	"todo-app/utils"
)

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserServiceImpl(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (s *UserServiceImpl) Register(request request.RegisterRequest) (response.UserResponse, error) {
	if err := utils.ValidatePassword(request.Password); err != nil {
		log.Printf("Password validation failed: %v", err)
		return response.UserResponse{}, err
	}
	user, err := s.userRepository.Register(request)
	if err != nil {
		return response.UserResponse{}, err
	}
	var userResponse response.UserResponse
	err = copier.Copy(&userResponse, &user)
	if err != nil {
		return userResponse, err
	}
	userResponse.Created = user.CreatedAt.Format(time.RFC3339)
	userResponse.Updated = user.UpdatedAt.Format(time.RFC3339)
	return userResponse, nil
}

func (s *UserServiceImpl) VerifyUser(username string, password string) (response.UserResponse, error) {
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		return response.UserResponse{}, errors.New("user not found")
	}
	err = utils.ComparePasswords(user.Password, password)
	if err != nil {
		return response.UserResponse{}, errors.New("invalid credentials")
	}
	var userResponse response.UserResponse
	err = copier.Copy(&userResponse, &user)
	if err != nil {
		return userResponse, err
	}

	return userResponse, nil
}

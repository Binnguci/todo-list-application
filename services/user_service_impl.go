package services

import (
	"github.com/jinzhu/copier"
	"time"
	"todo-app/dto/request"
	"todo-app/dto/response"
	"todo-app/repositories"
)

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserServiceImpl(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (s *UserServiceImpl) Register(request request.RegisterRequest) (response.UserResponse, error) {
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

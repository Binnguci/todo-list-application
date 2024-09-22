package services

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"log"
	"math/rand"
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

func (u *UserServiceImpl) Register(request request.RegisterRequest) (response.UserResponse, error) {
	err := u.validationRegister(request.Username, request.Email)
	if err != nil {
		log.Printf("Validation failed: %v", err)
		return response.UserResponse{}, err
	}
	if err := utils.ValidatePassword(request.Password); err != nil {
		log.Printf("Password validation failed: %v", err)
		return response.UserResponse{}, err
	}
	user, err := u.userRepository.Register(request)
	if err != nil {
		return response.UserResponse{}, err
	}
	var userResponse response.UserResponse
	err = copier.Copy(&userResponse, &user)
	if err != nil {
		return userResponse, err
	}
	userResponse.Role = "USER"
	userResponse.Created = user.CreatedAt.Format(time.RFC3339)
	userResponse.Updated = user.UpdatedAt.Format(time.RFC3339)
	return userResponse, nil
}

func (u *UserServiceImpl) VerifyUser(username string, password string) (response.UserResponse, error) {
	user, err := u.userRepository.FindByUsername(username)
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

func (u *UserServiceImpl) validationRegister(username string, email string) error {
	_, err := u.userRepository.FindByUsername(username)
	if err == nil {
		return errors.New("username already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("error checking username: %w", err)
	}

	_, err = u.userRepository.FindByEmail(email)
	if err == nil {
		return errors.New("email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("error checking email: %w", err)
	}
	return nil
}

func (u *UserServiceImpl) generateOTP() (string, error) {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"

	otp := make([]byte, 6)
	for i := range otp {
		otp[i] = digits[rand.Intn(len(digits))]
	}

	return string(otp), nil
}

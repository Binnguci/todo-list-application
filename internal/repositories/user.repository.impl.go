package repositories

import (
	"github.com/binnguci/todo-app/internal/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Register(user model.User) (model.User, error)
	FindByEmail(email string) bool
	FindByUsername(username string) bool
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) IUserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (uri *UserRepositoryImpl) Register(user model.User) (model.User, error) {
	err := uri.DB.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (uri *UserRepositoryImpl) FindByEmail(email string) bool {
	var user model.User
	row := uri.DB.Take(&user, "email = ?", email)
	return row.RowsAffected > 0
}

func (uri *UserRepositoryImpl) FindByUsername(username string) bool {
	var user model.User
	row := uri.DB.Take(&user, "username = ?", username)
	return row.RowsAffected > 0
}

package user

import (
	"gorm.io/gorm"
	"todo-app/domain/request"
	"todo-app/models"
	"todo-app/utils"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}

func (r *UserRepositoryImpl) Register(req request.RegisterRequest) (models.User, error) {
	if err := r.Db.Where("name = ?", "user").Error; err != nil {
		return models.User{}, err
	}
	hashpass, err := utils.HashPassword(req.Password)
	if err != nil {
		return models.User{}, err
	}
	user := models.User{
		Username: req.Username,
		Password: hashpass,
		Email:    req.Email,
		FullName: req.FullName,
		Status:   "off",
	}
	result := r.Db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (models.User, error) {
	var user models.User
	if err := r.Db.Where("username = ?", username).Preload("Roles").First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil

}

func (r *UserRepositoryImpl) FindByEmail(email string) (models.User, error) {
	var user models.User
	if err := r.Db.Where("email = ?", email).Preload("Roles").First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

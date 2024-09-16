package repositories

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"todo-app/dto/request"
	"todo-app/models"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}

func hashpass(password string) (encodePassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (r *UserRepositoryImpl) Register(req request.RegisterRequest) (models.User, error) {
	var defaultRole models.Role
	if err := r.Db.Where("name = ?", "user").First(&defaultRole).Error; err != nil {
		return models.User{}, err
	}
	hashpass, err := hashpass(req.Password)
	if err != nil {
		return models.User{}, err
	}
	user := models.User{
		Username: req.Username,
		Password: hashpass,
		Email:    req.Email,
		FullName: req.FullName,
		Status:   "active",
		Roles:    []models.Role{defaultRole},
	}
	result := r.Db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

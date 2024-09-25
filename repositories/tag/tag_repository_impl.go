package tag

import (
	"gorm.io/gorm"
	"todo-app/models"
)

type TagRepositoryImpl struct {
	DB *gorm.DB
}

func NewTagRepositoryImpl(db *gorm.DB) TagRepository {
	return &TagRepositoryImpl{DB: db}
}

func (t TagRepositoryImpl) FindAll() ([]models.Tag, error) {
	var tag []models.Tag
	if err := t.DB.Find(&tag).Error; err != nil {
		return nil, err
	}
	return tag, nil
}

func (t *TagRepositoryImpl) FindByID(id int) (*models.Tag, error) {
	var tag models.Tag
	if err := t.DB.First(&tag, id).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (t *TagRepositoryImpl) Create(tag models.Tag) (*models.Tag, error) {
	if err := t.DB.Create(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (t *TagRepositoryImpl) Update(id int, tag models.Tag) (*models.Tag, error) {
	var existingTag models.Tag
	if err := t.DB.First(&existingTag, id).Error; err != nil {
		return nil, err
	}
	existingTag.Name = tag.Name
	if err := t.DB.Save(&existingTag).Error; err != nil {
		return nil, err
	}
	return &existingTag, nil
}

func (t *TagRepositoryImpl) Delete(id int) error {
	if err := t.DB.Delete(&models.Tag{}, id).Error; err != nil {
		return err
	}
	return nil
}

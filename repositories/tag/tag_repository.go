package tag

import "todo-app/models"

type TagRepository interface {
	FindAll() ([]models.Tag, error)
	FindByID(id int) (*models.Tag, error)
	Create(tag models.Tag) (*models.Tag, error)
	Update(id int, tag models.Tag) (*models.Tag, error)
	Delete(id int) error
}

package tag

import "todo-app/models"

type TagService interface {
	FindAll() ([]models.Tag, error)
}

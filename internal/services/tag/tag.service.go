package tag

import (
	"todo-app/internal/models"
)

type TagService interface {
	FindAll() ([]models.Tag, error)
}

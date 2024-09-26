package tag

import (
	"todo-app/internal/models"
	"todo-app/internal/repositories/tag"
)

type TagServiceImpl struct {
	tagRepository tag.TagRepository
}

func NewTagServiceImpl(tagRepository tag.TagRepository) TagService {
	return &TagServiceImpl{tagRepository: tagRepository}
}

func (tsi *TagServiceImpl) FindAll() ([]models.Tag, error) {
	tag, err := tsi.tagRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tag, nil
}

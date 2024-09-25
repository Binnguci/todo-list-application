package tag

import (
	"todo-app/models"
	"todo-app/repositories/tag"
)

type TagServiceImpl struct {
	tagRepository tag.TagRepository
}

func NewTagServiceImpl(tagRepository tag.TagRepository) TagService {
	return &TagServiceImpl{tagRepository: tagRepository}
}

func (t *TagServiceImpl) FindAll() ([]models.Tag, error) {
	tag, err := t.tagRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tag, nil
}

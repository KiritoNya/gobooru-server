package interfaces

import "github.com/KiritoNya/gobooru-server/models"

type ITagRepository interface {
	GetTagByName(name string) (models.Tag, error)
	GetTagById(id int) (models.Tag, error)
}

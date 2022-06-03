package interfaces

import "github.com/KiritoNya/gobooru-server/models"

type ITagCategoryRepository interface {
	GetTagCategoryByName(name string) (models.TagCategory, error)
	GetTagCategiryById(id int) (models.TagCategory, error)
}

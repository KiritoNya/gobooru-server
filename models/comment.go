package models

import "github.com/jinzhu/gorm"

// Comment data model
type Comment struct {
	gorm.Model
	Version int
	PostId  Post
	Creator User
	Text    string
	Score   int
}

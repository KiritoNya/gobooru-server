package model

import "github.com/jinzhu/gorm"

// Comment data model
type Comment struct {
	gorm.Model
	Version   int
	PostID    int
	CreatorID int
	Creator   User
	Text      string
	Score     int
}

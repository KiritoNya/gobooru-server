package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Version            int
	Safety             PostSafety
	Source             string
	Type               PostType
	Checksum           string
	ChecksumMd5        string
	CanvasWidth        int
	CanvasHeight       int
	ContentUrl         string
	ThumbnailUrl       string
	Tags               []Tag  `gorm:"many2many:post_tags"`
	Relations          []Post `gorm:"many2many:post_relations"`
	CreatorID          int
	Creator            User
	Score              int
	TagCount           int
	FavouriteCount     int
	CommentCount       int
	RelationCount      int
	HasCustomThumbnail bool
	Comments           []Comment
	Pools              []Pool `gorm:"many2many:post_pools"`
}

type PostSafety string

const (
	Safe    PostSafety = "safe"
	Sketchy PostSafety = "sketchy"
	Unsafe  PostSafety = "unsafe"
)

type PostType string

const (
	Image     PostType = "image"
	Animation PostType = "animation"
	Video     PostType = "video"
	Flash     PostType = "flash"
	Youtube   PostType = "youtube"
)

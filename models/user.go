package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// User data model
type User struct {
	gorm.Model
	Username           string
	Email              string
	Rank               Rank
	LastLoginTime      time.Time
	AvatarStyle        AvatarStyle
	AvatarUrl          string
	CommentCount       int
	UploadedPostCount  int
	LikedPostCount     int
	DislikedPostCount  int
	FavouritePostCount int
	Version            int
}

// Rank defines the user permission
type Rank string

const (
	Restricted    Rank = "restricted"
	Regular       Rank = "regular"
	Power         Rank = "power"
	Moderator     Rank = "moderator"
	Administrator Rank = "administrator"
)

// AvatarStyle define the avatar style
type AvatarStyle string

const (
	// Gravatar is an avatar image generated from email
	Gravatar AvatarStyle = "gravatar"

	// Manual is an avatar uploaded by user
	Manual AvatarStyle = "manual"
)

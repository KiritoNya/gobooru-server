package main

import (
	"context"
	"github.com/KiritoNya/gobooru-server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type RecorderLogger struct {
	logger.Interface
	Statements []string
}

func (r *RecorderLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, _ := fc()
	r.Statements = append(r.Statements, sql)
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	recorder := RecorderLogger{logger.Default.LogMode(logger.Info), nil}
	session := db.Session(&gorm.Session{
		Logger: &recorder,
	})
	err = session.AutoMigrate(
		&models.TagCategory{},
		&models.Tag{},
		&models.User{},
		&models.Comment{},
		&models.PoolCategory{},
		&models.Pool{},
		&models.Post{},
	)
	if err != nil {
		panic(err)
	}
}

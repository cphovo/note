package model

import (
	"time"

	"gorm.io/gorm"
)

type Aritcle struct {
	gorm.Model
	Title    string
	Content  string
	FileInfo FileInfo
	Tags     []Tag `gorm:"many2many:article_tags;"`
}

type Tag struct {
	gorm.Model
	Name string
}

type FileInfo struct {
	gorm.Model
	AritcleID     uint
	FileName      string
	FileCreatedAt time.Time
}

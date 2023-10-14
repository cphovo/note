package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title      string `gorm:"comment:标题"`
	Content    string `gorm:"comment:内容"`
	FileName   string `gorm:"comment:文件名称"`
	FileTypeID uint   `gorm:"comment:文件类型"`
	FileType   FileType
	Date       string `gorm:"comment:日期"`
	Tags       []Tag  `gorm:"many2many:article_tags;comment:标签"`
}

type Tag struct {
	gorm.Model
	Name string `gorm:"unique;not null;comment:标签名称"`
}

type FileType struct {
	gorm.Model
	TypeName   string `gorm:"unique;not null;comment:文件类型名称"`
	Ext        string `gorm:"not null;comment:文件扩展名"`
	Supported  bool   `gorm:"comment:是否支持"`
	Rendered   bool   `gorm:"comment:是否渲染"`
	Compressed bool   `gorm:"comment:是否压缩"`
	Articles   []Article
}

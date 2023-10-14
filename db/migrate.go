package db

import (
	"log"

	"github.com/cphovo/note/model"
)

var fileTypes = []model.FileType{
	{
		TypeName:   "markdown",
		Ext:        ".md",
		Supported:  true,
		Rendered:   true,
		Compressed: false,
	},
	{
		TypeName:   "txt",
		Ext:        ".txt",
		Supported:  true,
		Rendered:   true,
		Compressed: false,
	},
	{
		TypeName:   "pdf",
		Ext:        ".pdf",
		Supported:  true,
		Rendered:   false,
		Compressed: true,
	},
	{
		TypeName:   "docx",
		Ext:        ".docx",
		Supported:  true,
		Rendered:   false,
		Compressed: true,
	},
	{
		TypeName:   "go",
		Ext:        ".go",
		Supported:  true,
		Rendered:   true,
		Compressed: false,
	},
}

func MigrateFileType() {
	d, err := GetDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	for _, fileType := range fileTypes {
		d.DB.FirstOrCreate(&fileType, model.FileType{TypeName: fileType.TypeName})
	}
}

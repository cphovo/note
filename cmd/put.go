package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/cphovo/note/db"
	"github.com/cphovo/note/model"
	"github.com/cphovo/note/utils"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put [filepath]",
	Short: "Archiving your file",
	Long: `Archiving your file provides a way to safely store your important files. 

Using this command, you can specify the filepath of the file you want to archive. This ensures your files are backed up and protected.
	
For example:
To archive a file named 'document.txt' located in the current directory, you would use:
$ note put document.txt
	
This will store 'document.txt' in the archive.`,
	Run: func(cmd *cobra.Command, args []string) {
		save(args) // save file info into database
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}

func save(filepaths []string) {
	if len(filepaths) == 0 {
		log.Fatalf("Please specify a filepath")
	}

	d, err := db.GetDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	tx := d.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, filepath := range filepaths {
		fi, err := utils.GetFileInfo(filepath)
		if err != nil {
			tx.Rollback()
			log.Fatal("Could not get file info: ", err)
		}

		ext := fi.Ext
		if ext == "" {
			ext = ".txt"
		}

		var fileType model.FileType
		if err := d.DB.Where(&model.FileType{Ext: ext}).First(&fileType).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Fatalf("Unsupported file type: %v", fi.Ext)
			}
		}

		content := fi.Content

		desc, ok := utils.CheckIfContainsDescription(fi.Content)
		if !ok {
			desc = utils.Description{
				Describe: strings.TrimSuffix(fi.Name, fi.Ext),
				Date:     fi.CreatedAt.Format("2006/01/02"),
			}
			if ext == ".md" {
				content = fmt.Sprintf("%s\n\n%s", desc.Code(), fi.Content)
			}
		}

		title := desc.Describe
		if len(title) >= 20 {
			title = title[:20]
		}

		aritcle := model.Article{
			Title:      title,
			Content:    content,
			FileName:   fi.Name,
			FileTypeID: fileType.ID,
			Date:       desc.Date,
		}

		if err := tx.Create(&aritcle).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Failed to insert article: %v", err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}
}

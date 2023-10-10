package cmd

import (
	"fmt"
	"strings"

	"github.com/cphovo/note/constants"
	"github.com/cphovo/note/db"
	"github.com/cphovo/note/model"
	"github.com/cphovo/note/utils"
	"github.com/spf13/cobra"
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
		if len(args) == 0 {
			fmt.Println("Please specify a filepath")
			return
		}
		var fileInfos []utils.FileInfo
		for _, filepath := range args {
			fi, err := utils.GetFileInfo(filepath)
			if err != nil {
				fmt.Println(err)
				return
			}
			fileInfos = append(fileInfos, *fi)
		}
		fmt.Println(fileInfos)
		database, err := db.NewDatabase(constants.DB_CONNECTION_STRING)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer database.Close()
		for _, fileInfo := range fileInfos {
			names := strings.Split(fileInfo.Name, ".")
			aritcle := model.Aritcle{
				Title:   strings.Join(names[:len(names)-1], "."),
				Content: fileInfo.Content,
				FileInfo: model.FileInfo{
					FileName:      fileInfo.Name,
					FileCreatedAt: fileInfo.CreatedAt,
				},
			}
			database.DB.Create(&aritcle)
		}
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}

package utils

import (
	"io"
	"os"
	"regexp"
	"time"
)

type FileInfo struct {
	Name      string
	Content   string
	CreatedAt time.Time
}

// Function to get the file info of a given filepath
func GetFileInfo(filepath string) (*FileInfo, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return &FileInfo{
		Name:      fileInfo.Name(),
		Content:   string(bytes),
		CreatedAt: fileInfo.ModTime(),
	}, nil
}

// Remove code blocks in MarkDown file content
func RemoveCodeBlocks(content string) string {
	re := regexp.MustCompile("```.*?\n.*?```")
	return re.ReplaceAllString(content, "")
}

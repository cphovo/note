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

// This function takes a string as an argument and returns a string with code blocks removed
func RemoveCodeBlocks(content string) string {
	// Compile a regular expression to match code blocks, excluding plain text block
	re := regexp.MustCompile("(?s)```\\w+\\n.*?```")
	return re.ReplaceAllString(content, "")
}

// Remove empty lines from a given string
func RemoveEmptyLines(content string) string {
	re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)
	return re.ReplaceAllString(content, "")
}

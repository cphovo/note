package utils

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

type FileInfo struct {
	Name      string
	Content   string
	CreatedAt time.Time
}

type Config struct {
	// the path of the sqlite db saved.
	Path string `json:"path"`
	// the name of the sqlite db
	Name string `json:"name"`
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

// Create a directory if it does not exist
func MkdirIfNotExist(path string) (string, error) {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return "", errors.New("error creating directory")
		}
	}
	return dir, nil
}

func HomeDir() string {
	s, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return s
}

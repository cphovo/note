package utils

import (
	"errors"
	"fmt"
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
	Ext       string
	CreatedAt time.Time
}

// Function to get the file info of a given filepath
func GetFileInfo(path string) (*FileInfo, error) {
	file, err := os.Open(path)
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
		Ext:       filepath.Ext(path),
		CreatedAt: fileInfo.ModTime(),
	}, nil
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

// CopyFile copies a file from src to dst.
func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	// Check if the source is a regular file
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
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

// Return the first code block found in the string.
func FindFirstCodeBlock(content string) string {
	re := regexp.MustCompile("(?s)```(\\w+)?\\n(.*?)\\n```")
	matches := re.FindStringSubmatch(content)

	if len(matches) > 0 {
		return matches[0]
	}

	return ""
}

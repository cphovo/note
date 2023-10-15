package config

import (
	"encoding/json"
	"fmt"
	"go/build"
	"log"
	"os"
	"path/filepath"

	"github.com/cphovo/note/utils"
)

const (
	DbFileName     = "data.db"
	ConfigDirName  = ".note"
	ConfigFileName = "config.json"
	JieBaDirName   = "jieba"
)

type Config struct {
	// the path of the sqlite db saved.
	Path string `json:"path"`

	// the dir path of the default jieba dataset saved
	JieBa utils.JieBa `json:"jieba"`
}

var (
	loadedConfig *Config
	configPath   = filepath.Join(utils.HomeDir(), ConfigDirName, ConfigFileName)
	jieBaDirPath = filepath.Join(utils.HomeDir(), ConfigDirName, JieBaDirName)
)

func init() {
	var err error
	loadedConfig, err = loadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
}

func GlobalConfig() *Config {
	return loadedConfig
}

func loadConfig() (*Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return saveDefaultConfig()
	}

	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var conf Config
	err = json.Unmarshal(b, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func saveDefaultConfig() (*Config, error) {
	path, err := utils.MkdirIfNotExist(configPath)
	if err != nil {
		return nil, err
	}

	defaultConfig := &Config{
		Path: filepath.Join(path, DbFileName),
		JieBa: utils.JieBa{
			DictPath:      filepath.Join(path, JieBaDirName, "jieba.dict.utf8"),
			HmmPath:       filepath.Join(path, JieBaDirName, "hmm_model.utf8"),
			UserDictPath:  filepath.Join(path, JieBaDirName, "user.dict.utf8"),
			IdfPath:       filepath.Join(path, JieBaDirName, "idf.utf8"),
			StopWordsPath: filepath.Join(path, JieBaDirName, "stop_words.utf8"),
		},
	}

	b, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(configPath, b, 0644); err != nil {
		return nil, err
	}

	if err := copyJieBaDefaultDataSet(jieBaDirPath); err != nil {
		return nil, err
	}

	fmt.Println("Initialize configuration file successfully.")

	return defaultConfig, nil
}

func copyJieBaDefaultDataSet(dstDir string) error {
	pkg, err := build.Import("github.com/yanyiwu/gojieba", "", build.FindOnly)
	if err != nil {
		return err
	}

	dictDir := filepath.Join(pkg.Dir, "dict")
	os.MkdirAll(dstDir, os.ModePerm)

	files := []string{
		"jieba.dict.utf8",
		"hmm_model.utf8",
		"user.dict.utf8",
		"idf.utf8",
		"stop_words.utf8",
	}

	for _, file := range files {
		srcPath := filepath.Join(dictDir, file)
		dstPath := filepath.Join(dstDir, file)
		if _, err := utils.CopyFile(srcPath, dstPath); err != nil {
			return err
		}
	}

	return nil
}

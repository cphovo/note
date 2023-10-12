package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/cphovo/note/utils"
)

const (
	DbFileName     = "data.db"
	ConfigDirName  = ".note"
	ConfigFileName = "config.json"
)

type Config struct {
	// the path of the sqlite db saved.
	Path string `json:"path"`
}

var (
	loadedConfig *Config
	configPath   = filepath.Join(utils.HomeDir(), ConfigDirName, ConfigFileName)
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
	}
	b, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(configPath, b, 0644)
	if err != nil {
		return nil, err
	}

	return defaultConfig, nil
}

package db

import (
	"os"
	"sync"

	"github.com/cphovo/note/constants"
	"github.com/cphovo/note/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

var (
	instance *Database
	once     sync.Once
)

// GetDB is a function that returns a Database instance and an error
func GetDB() (*Database, error) {
	dsn := constants.DB_PATH

	envDsn, ok := os.LookupEnv("DB_PATH")
	if ok {
		dsn = envDsn
	}

	var err error
	// Execute the function only once
	once.Do(func() {
		instance, err = initDB(dsn)
	})
	return instance, err
}

func initDB(dsn string) (*Database, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	for _, model := range model.Models() {
		err = db.AutoMigrate(model)
		if err != nil {
			return nil, err
		}
	}
	return &Database{DB: db}, nil
}

func (database *Database) Close() error {
	db, err := database.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

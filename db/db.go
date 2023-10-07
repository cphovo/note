package db

import (
	"os"

	"github.com/cphovo/note/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func InitDB(dsn string) error {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	defer func() {
		d, err := db.DB()
		if err != nil {
			panic(err)
		}
		d.Close()
	}()

	for _, model := range model.Models() {
		err = db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}
	return nil
}

func InitDBIfNotExist(dsn string) error {
	if _, err := os.Stat(dsn); os.IsNotExist(err) {
		file, err := os.Create(dsn)
		if err != nil {
			return err
		}
		file.Close()
	}
	return InitDB(dsn)
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
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

package config

import (
	"os"

	"github.com/SantGT5/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	dbPath := "./db/main.db"
	logger := GetLogger("sqlite")

	// Check if db file exist
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("db file not found, creating...")

		// Create db file and directory
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite opening error: %v", err.Error())
		return nil, err
	}

	// Migrate Schema
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("sqlite AutoMigration error: %v", err.Error())
		return nil, err
	}

	return db, nil
}

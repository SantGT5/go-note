package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// init SQLite
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error init sqlite: %v", err.Error())
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// init Logger
	logger = NewLogger(p)

	return logger
}

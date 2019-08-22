package db

import (
	"github.com/jinzhu/gorm"

	// We are using postgres as database
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/protonhq/proton/domain"
)

// InitDB - Create Database Connection
func InitDB(connection string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connection)
	return db, err
}

// Migrate - Migrate Database
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&domain.Account{})
}

package database

import (
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// FileName from the sqlite database file.
const FileName = "perdat.db"

// Connector is responsible to provide database connection.
type Connector interface {
	Connect(string) (*gorm.DB, error)
}

// Connect to a sqlite database.
func Connect() (*gorm.DB, error) {
	logr, err := NewFileLogger()
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(
		sqlite.Open(FileName), &gorm.Config{Logger: logr})
	if err != nil {
		return nil, errors.Wrap(err, "DB connection")
	}

	return db, nil
}

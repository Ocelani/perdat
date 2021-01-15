package database

import (
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connector is responsible to provide database connection.
type Connector interface {
	Connect(string) (*DB, error)
}

// DB is the database provider type.
type DB struct {
	Conn *gorm.DB
	File string
}

// Connect to a sqlite database.
func Connect(file string, logr logger.Interface) (*DB, error) {
	db, err := gorm.Open(
		sqlite.Open(file), &gorm.Config{Logger: logr})

	if err != nil {
		return nil, errors.Wrap(err, "db connect")
	}

	return &DB{db, file}, nil
}

package database

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3" // sqlite db
	"github.com/pkg/errors"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

// Connector is responsible to provide database connection.
type Connector interface {
	Connect(string) (*DB, error)
}

// DB is the database provider type.
type DB struct {
	Engine *xorm.Engine
	File   string
}

// InstantiateLogger instantiates a logger to register database operations.
func (db *DB) InstantiateLogger() error {
	l, err := os.Create("sql.log")
	if err != nil {
		return errors.Wrap(err, "db logger")
	}
	db.Engine.SetLogger(log.NewSimpleLogger(l))
	return nil
}

// Connect is responsible to make a connection with sqlite database.
func Connect(file string) (*DB, error) {
	engine, err := xorm.NewEngine("sqlite3", fmt.Sprintf("./database/%s", file))
	if err != nil {
		return nil, errors.Wrap(err, "db connect")
	}
	engine.SetMapper(names.GonicMapper{})

	db := &DB{
		Engine: engine,
		File:   file,
	}
	return db, nil
}

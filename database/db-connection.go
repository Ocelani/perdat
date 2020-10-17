package database

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3" // sqlite db
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

// Connect is responsible to make a connection with sqlite database.
func Connect(db string) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("sqlite3", fmt.Sprintf("./%s", db))
	engine.SetMapper(names.GonicMapper{})
	f, err := os.Create("sql.log")

	if err != nil {
		return nil, err
	}
	engine.SetLogger(log.NewSimpleLogger(f))

	return engine, nil
}

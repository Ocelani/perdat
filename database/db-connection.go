package database

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3" // sqlite db
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

func Connect() *xorm.Engine {
	engine, err := xorm.NewEngine("sqlite3", fmt.Sprintf("./perdat.db"))
	engine.SetMapper(names.GonicMapper{})

	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(log.NewSimpleLogger(f))

	return
}

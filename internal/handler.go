package internal

import (
	"fmt"

	"github.com/Ocelani/perdat/pkg/common/database"
	"github.com/Ocelani/perdat/pkg/common/logger"
	"github.com/Ocelani/perdat/pkg/fact"
	"github.com/Ocelani/perdat/pkg/habit"
)

const (
	_logFile = "./logs/sql.log"
	_dbFile  = "./perdat.db"
)

func FactHandler(operation string, args []string) {
	var (
		logr = logger.NewLogger()
		db   = database.SQLiteOpen(_dbFile, database.NewFileLogger(_logFile))
	)
	facts, err := fact.Handler(db, operation, args)
	if err != nil {
		logr.Error().Err(err).Send()
		return
	}
	for _, f := range *facts {
		logr.FactCompleteLog(operation, &f)
	}
}

func HabitHandler(operation string, args []string) {
	var (
		logr = logger.NewLogger()
		db   = database.SQLiteOpen(_dbFile, database.NewFileLogger(_logFile))
	)
	habits, err := habit.Handler(db, operation, args)
	if err != nil {
		logr.Error().Err(err).Send()
		return
	}
	for _, h := range *habits {
		fmt.Println(h)
	}
}

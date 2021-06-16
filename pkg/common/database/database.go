package database

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/Ocelani/perdat/pkg/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct{ *gorm.DB }

// SQLiteOpen just opens the sqlite file database.
func SQLiteOpen(file string, logr logger.Interface) *DB {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{Logger: logr})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&entity.Fact{}, &entity.Counter{}, &entity.Habit{}); err != nil {
		panic(err)
	}
	return &DB{db}
}

// NewStdoutLogger instantiates a stdout logger to print the output of database operations.
func NewStdoutLogger() logger.Interface {
	return newLogger(os.Stdout)
}

// NewFileLogger instantiates a text file logger to register database operations.
func NewFileLogger(filePath string) logger.Interface {
	if filePath == "" {
		return nil
	}
	f, _ := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer f.Close()
	
	return newLogger(f)
}

func newLogger(wr io.Writer) logger.Interface {
	return logger.New(
		log.New(wr, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Millisecond * 500, // Slow SQL threshold
			LogLevel:      logger.Silent,          // Log level
			Colorful:      true,                   // Disable color
		},
	)
}

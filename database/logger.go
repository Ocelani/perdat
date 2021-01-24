package database

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm/logger"
)

// NewStdoutLogger instantiates a stdout logger to print the output of database operations.
func NewStdoutLogger() logger.Interface {
	return newLogger(os.Stdout)
}

// NewFileLogger instantiates a text file logger to register database operations.
func NewFileLogger() (logger.Interface, error) {
	f, err := os.OpenFile(
		"./internal/logs/sql.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return nil, errors.Wrap(err, "DB file logger")
	}
	return newLogger(f), nil
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

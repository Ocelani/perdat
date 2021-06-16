package common

import (
	"fmt"

	"github.com/Ocelani/perdat/pkg/common/database"
)

type Repository struct {
	DB  *database.DB
	Err error
}

func NewRepository(db *database.DB) *Repository {
	return &Repository{DB: db}
}

func NewSQLiteRepository(dbFilePath, logFilePath string) *Repository {
	return &Repository{
		DB: database.SQLiteOpen(dbFilePath, database.NewFileLogger(logFilePath)),
	}
}

func GetDB(repository interface{}) (*database.DB, error) {
	r, ok := repository.(*Repository)
	if !ok {
		return nil, fmt.Errorf("GetDB failed: not a repository")
	}
	return r.DB, nil
}

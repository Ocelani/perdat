package logger

import (
	"os"

	"github.com/Ocelani/perdat/pkg/entity"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Log struct{ zerolog.Logger }

func NewLogger() *Log {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	return &Log{log.Output(zerolog.ConsoleWriter{Out: os.Stdout})}
}

func (l *Log) FactBasicLog(op string, t *entity.Fact) {
	l.Info().
		Uint("ID", t.ID).
		Str("name", t.Name).
		Msg(op)
}

func (l *Log) FactCompleteLog(op string, t *entity.Fact) {
	l.Info().
		Uint("ID", t.ID).
		Str("name", t.Name).
		// Bool("done", t.Done).
		Str("createdAt", t.CreatedAt.Format("2006-01-02")).
		Str("updatedAt", t.UpdatedAt.Format("2006-01-02")).
		Msg(op)
}

func (l *Log) FactDoneInfoLog(op string, t *entity.Fact) {
	l.Info().
		Uint("ID", t.ID).
		Str("name", t.Name).
		// Bool("done", t.Done).
		Msg(op)
}

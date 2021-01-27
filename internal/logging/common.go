package logging

import (
	"io"

	"github.com/rs/zerolog"
)

const (
	Debug       = "DEBUG"
	Application = "APPLICATION"
)

type FormattedLog map[string]string

type Log struct {
	logger   *zerolog.Logger
	logLevel string
}

func New(output io.Writer, logLevel string) (*Log, error) {
	// zerolog.ConsoleWriter{Out: os.Stderr}
	logger := zerolog.New(output).With().Timestamp().Logger()

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if logLevel != Debug {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	return &Log{
		logger:   &logger,
		logLevel: logLevel,
	}, nil
}

func (l *Log) LogDebugAndApplication(params FormattedLog) {
	l.LogApplication(params)
	l.LogDebug(params)
}

func (l *Log) LogDebugAndError(params FormattedLog) {
	l.LogError(params)
	l.LogDebug(params)
}

func (l *Log) LogFatal(params FormattedLog) {
	fatal := l.logger.Fatal()
	for k, v := range params {
		fatal.Str(k, v)
	}
	fatal.Send()
}

func (l *Log) LogDebug(params FormattedLog) {
	debug := l.logger.Debug()
	for k, v := range params {
		debug.Str(k, v)
	}
	debug.Send()
}

func (l *Log) LogApplication(params FormattedLog) {
	app := l.logger.Info()
	for k, v := range params {
		app.Str(k, v)
	}
	app.Send()
}

func (l *Log) LogError(params FormattedLog) {
	err := l.logger.Error()
	for k, v := range params {
		err.Str(k, v)
	}
	err.Send()
}

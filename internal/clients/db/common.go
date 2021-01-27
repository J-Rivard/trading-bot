package db

import (
	"database/sql"
	"fmt"

	"github.com/J-Rivard/trading-bot/internal/logging"
	_ "github.com/lib/pq"
)

type Parameters struct {
	Username string
	Password string
	Host     string
	Schema   string
	DBName   string
}

type DB struct {
	Client *sql.DB
	Events []*string
	Log    *logging.Log
}

func New(params *Parameters, log *logging.Log) (*DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s database=%s search_path=%s sslmode=disable",
		params.Username, params.Password, params.Host, params.DBName, params.Schema)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: db,
		Log:    log,
	}, nil
}

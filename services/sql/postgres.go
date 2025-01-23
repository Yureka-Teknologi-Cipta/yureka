package yureka_sql

import (
	"context"
	"time"

	_ "github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
)

func ConnectPostgres(timeout time.Duration, dbURL string) *sqlx.DB {
	_, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	db, err := sqlx.Open("pgx", dbURL)
	if err != nil {
		panic("Cannot connect to database: " + err.Error())
	}

	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(2 * time.Minute)
	db.SetMaxOpenConns(40)

	if err = db.Ping(); err != nil {
		panic("Database not reachable: " + err.Error())
	}

	return db
}

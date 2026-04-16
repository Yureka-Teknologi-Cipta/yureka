package yureka_sql

import (
	"context"
	"net/url"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

// ConnectMSSQL connects to SQL Server database using sqlx and returns the connection.
// It takes a timeout duration and a database URL as parameters.
// The database URL should be in the format: sqlserver://username:password@host:port/database?query_params
//
// Example: sqlserver://sa:password@localhost:1433/go-template?encrypt=disable
func ConnectMSSQL(timeout time.Duration, dbURL string) *sqlx.DB {
	_, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// convert dbURL to DSN
	parsedURL, err := url.Parse(dbURL)
	if err != nil {
		panic("Invalid database URL: " + err.Error())
	}

	// database in sqlserver is set in query param, not path, so
	query := parsedURL.Query()
	if dbName := query.Get("database"); dbName == "" {
		query.Set("database", parsedURL.Path[1:]) // remove leading slash
		// remove path since sqlserver driver doesn't use it
		parsedURL.Path = ""
	}

	// restore query to raw query
	parsedURL.RawQuery = query.Encode()

	db, err := sqlx.Open("sqlserver", parsedURL.String())
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

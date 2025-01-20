package yureka_sql

import (
	"context"
	"fmt"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectMysql(timeout time.Duration, dbURL string) *sqlx.DB {
	_, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// convert dbURL to DSN
	parsedURL, err := url.Parse(dbURL)
	if err != nil {
		panic("Invalid database URL: " + err.Error())
	}

	user := parsedURL.User.Username()
	password, hasPassword := parsedURL.User.Password()
	host := parsedURL.Host
	dbName := parsedURL.Path
	query := parsedURL.Query().Encode()

	var dsn string
	if hasPassword {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)%s?%s", user, password, host, dbName, query)
	} else {
		dsn = fmt.Sprintf("%s@tcp(%s)%s?%s", user, host, dbName, query)
	}

	db, err := sqlx.Open("mysql", dsn)
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

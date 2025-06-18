package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql" // the _ because not using it in here, but the db needs the driver so must import
)

const (
	maxOpenDBConn = 25
	maxIdleDBConn = 25
	maxDBLifetime = 5 * time.Minute
)

func initMySQLDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// test our database
	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetMaxIdleConns(maxIdleDBConn)

	return db, nil
}

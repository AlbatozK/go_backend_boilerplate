package db

import (
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	pool *sqlx.DB
	err  error
)

func Init() error {
	conn_string := os.Getenv("PG_CONN_STRING")
	pool, err = sqlx.Connect("postgres", conn_string)
	if err != nil {
		return err
	}
	pool.SetMaxIdleConns(5)
	pool.SetConnMaxLifetime(2 * time.Minute)
	pool.SetMaxOpenConns(95)
	return nil
}

func GetPool() *sqlx.DB {
	return pool
}

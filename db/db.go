package db

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	pool *sqlx.DB
	err  error
)

func Init() {
	conn_string := os.Getenv("PG_CONN_STRING")
	pool, err = sqlx.Connect("postgres", conn_string)
	pool.SetMaxIdleConns(5)
	pool.SetConnMaxLifetime(2 * time.Minute)
	pool.SetMaxOpenConns(95)
	if err != nil {
		log.Println("m=GetPool,msg=connection has failed", err)
	}
}

func GetPool() *sqlx.DB {
	return pool
}

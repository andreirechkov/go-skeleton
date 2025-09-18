package db

import (
	"database/sql"
	"fmt"
	"os"
	
	_ "github.com/lib/pq"
)

func NewPostgres() (*sql.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, pass, dbname,
    )
    return sql.Open("postgres", connStr)
}
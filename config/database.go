package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	return sql.Open("postgres", os.Getenv("DATABASE_URL"))
}

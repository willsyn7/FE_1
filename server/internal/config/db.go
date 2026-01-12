package config

import (
	"database/sql" // https://pkg.go.dev/database/sql
	// defeault library
	// used for os and to lookup env file
	_ "github.com/lib/pq" // PostgreSQL driver: https://pkg.go.dev/github.com/lib/pq
)

var DB *sql.DB

func InitDb() error {

}

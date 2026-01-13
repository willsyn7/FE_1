package config

import (
	//stanrd lib
	"context"
	"fmt"
	"os"
	// "time"
	"log"

	//other
	"github.com/jackc/pgx/v5"
	
//internal
	// "github.com/jackc/pgx/v5/pgxpool"
	
)

// DB is the single database connection
var DB *pgx.Conn

func InitDb() error {
	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		return fmt.Errorf("DATABASE_URL is missing from env")
	}

	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Test the connection
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		conn.Close(context.Background())
		return fmt.Errorf("query failed: %w", err)
	}

	// Assign to global DB variable so you can use it throughout your app
	DB = conn
	log.Println("Connected to database:", version)
	return nil
}

/*
import "ToDoGo/internal/config"

// Query single row
var name string
err := config.DB.QueryRow(ctx, "SELECT name FROM users WHERE id = $1", userID).Scan(&name)

// Query multiple rows
rows, err := config.DB.Query(ctx, "SELECT id, title FROM todos")
defer rows.Close()

for rows.Next() {
    var id int
    var title string
    rows.Scan(&id, &title)
}

// Insert/Update/Delete
_, err := config.DB.Exec(ctx, "INSERT INTO todos (title) VALUES ($1)", "New todo")
*/
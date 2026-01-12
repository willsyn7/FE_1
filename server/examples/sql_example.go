package examples

import (
	"ToDoGo/internal/config/db"
)

// Example: How to write SQL queries using backticks in Go
// In Go, you use backticks (`) for raw string literals, similar to template literals in Node.js

func ExampleSQLQueries() {
	// Single-line SQL query
	query1 := `SELECT * FROM todos WHERE id = $1`

	// Multi-line SQL query (this is where backticks really shine!)
	query2 := `
		SELECT 
			id,
			title,
			description,
			completed,
			created_at
		FROM todos
		WHERE completed = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	// Complex query with joins
	query3 := `
		SELECT 
			t.id,
			t.title,
			t.description,
			u.username as created_by
		FROM todos t
		INNER JOIN users u ON t.user_id = u.id
		WHERE t.completed = $1
		AND u.active = true
	`

	// Using the queries (example)
	_ = query1
	_ = query2
	_ = query3

	// Example: Execute a query
	rows, err := db.DB.Query(query2, false, 10)
	if err != nil {
		// handle error
		return
	}
	defer rows.Close()

	// Example: Execute with parameters
	_, err = db.DB.Exec(`
		INSERT INTO todos (title, description, completed)
		VALUES ($1, $2, $3)
	`, "My Todo", "Description here", false)

	// Example: Prepared statement
	stmt := `
		UPDATE todos
		SET completed = $1, updated_at = NOW()
		WHERE id = $2
	`
	_, err = db.DB.Exec(stmt, true, 123)
}

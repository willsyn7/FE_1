package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"ToDoGo/internal/config"
	"ToDoGo/internal/routes"
)

func main() {
	const port = ":8080"

	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Database
	err = config.InitDb()
	if err != nil {
		log.Printf("Warning: Database connection failed: %v\n", err)
		log.Println("Server will start without database connection")
	} else {
		log.Println("Database connected successfully")
	}

	// Routes
	r := routes.SetupRouter()



	// Start server
	fmt.Printf("server is running on port %s\n", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Listen and Server", err)
	}
}
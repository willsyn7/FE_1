package main

import (
	//pacakges
	"fmt"
	"log"
	"net/http"
		"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	// "os"


	//interneal
	"ToDoGo/internal/config"
	



)

type Response struct {
}

func main() {
	const port = ":8080"

	r := mux.NewRouter()

	r.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Test1")
	}).Methods("GET")


	

	// err := http.ListenAndServe(port,nil)
	//
	//	if err != nil {
	//		log.Fatal("ListenandServe", err)

	//	}
	//load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}


//db	

	err = config.InitDb()
	if err != nil{
		log.Printf("Warning: Database connection failed: %v\n", err)
		log.Println("Server will start without database connection")
	} else {
		log.Println("Database connected successfully")
	}


//port
	fmt.Printf("server is running on port %s\n", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Listen and Server", err)
	}

fmt.Printf(`hi`)


	//default Handler
	// http.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Test1")
	// })

	// fmt.Printf("server is running on port %s\n", port)
	// err := http.ListenAndServe(port, nil) // nil = use default handler
	// if err != nil {
	// 	log.Fatal("Listen and Server", err)
	// }
}

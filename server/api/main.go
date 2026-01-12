package main

import (
	"fmt"
	"log"
	"net/http"

	// "os"
	// "ToDoGo/internal/config"
	// "github.com/joho/godotenv"
	"github.com/gorilla/mux"
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

	fmt.Printf("server is running on port %s\n", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Listen and Server", err)
	}

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

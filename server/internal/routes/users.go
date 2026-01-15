package routes

import (
	"github.com/gorilla/mux"

	"ToDoGo/internal/handlers"
)

func RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/signup", handlers.SignUp).Methods("POST") // register rotuers
	r.HandleFunc("/delete" , handlers.DeleteUser).Methods("POST") // Delete 
}
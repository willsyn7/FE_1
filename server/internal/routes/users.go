package routes

import (
	"github.com/gorilla/mux"

	"ToDoGo/internal/handlers"
)

func RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/user/signup", handlers.SignUp).Methods("POST") // register rotuers
	r.HandleFunc("/user/delete" , handlers.DeleteUser).Methods("POST") // Delete 
	r.HandleFunc("/user/select", handlers.GetUserData).Methods("GET") // SelectUser
}
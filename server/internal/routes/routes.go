package routes

import "github.com/gorilla/mux"

func SetupRouter() *mux.Router {
	r := mux.NewRouter() // Intiatilzie new router
	RegisterUserRoutes(r) // register user routers to r whcihis router and return r back to call site whcih woudl be in main.go
	return r
}

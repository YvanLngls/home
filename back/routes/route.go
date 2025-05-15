package routes

import (
	"back/handlers"
	"back/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/protected", middleware.AuthMiddleware(handlers.ProtectedHandler)).Methods("GET")

	return r
}

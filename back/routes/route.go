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
	r.HandleFunc("/ws", handlers.WebSocketHandler)
	r.HandleFunc("/verify-token", middleware.AuthMiddleware(handlers.VerifyTokenHandler)).Methods("POST")
	r.HandleFunc("/car/add", middleware.AuthMiddleware(handlers.AddCarHandler)).Methods("POST")
	r.HandleFunc("/car/delete", middleware.AuthMiddleware(handlers.DeleteCarHandler)).Methods("DELETE")
	r.HandleFunc("/car/get", middleware.AuthMiddleware(handlers.GetCarHandler)).Methods("GET")

	return r
}

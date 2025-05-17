package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/rs/cors"
	"back/config"
	"back/routes"
)

func main() {
	config.InitRedis()

	r := routes.SetupRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Nuxt dev server
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
	})

	handler := c.Handler(r)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"back/config"
	"back/routes"
)

func main() {
	config.InitRedis()

	r := routes.SetupRoutes()

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

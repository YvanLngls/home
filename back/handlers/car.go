package handlers

import (
	"context"
	"back/config"
	"back/models"
	"encoding/json"
	"fmt"
	"net/http"
)


func AddCarHandler(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the car using Redis INCR
	carID, err := config.RedisDataDb.Incr(context.Background(), "car:id").Result()
	if err != nil {
		http.Error(w, "Failed to generate car ID", http.StatusInternalServerError)
		return
	}

	// Set the ID in the car object
	car.ID = fmt.Sprintf("%d", carID)

	// Marshal the car object to JSON
	carJSON, err := json.Marshal(car)
	if err != nil {
		http.Error(w, "Failed to marshal car data", http.StatusInternalServerError)
		return
	}

	// Store the car object in Redis with the generated ID as the key
	err = config.RedisDataDb.Set(context.Background(), "car:"+car.ID, carJSON, 0).Err()
	if err != nil {
		http.Error(w, "Failed to store car", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Car added successfully with ID: %s", car.ID)
	fmt.Println("Car added successfully with ID:", car.ID)
}

func DeleteCarHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}

	key := "car:" + id

	// Check if the car exists
	exists, err := config.RedisDataDb.Exists(context.Background(), key).Result()
	if err != nil {
		http.Error(w, "Failed to check if car exists", http.StatusInternalServerError)
		return
	}
	if exists == 0 {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	// Delete the car
	err = config.RedisDataDb.Del(context.Background(), key).Err()
	if err != nil {
		http.Error(w, "Failed to delete car", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Car deleted successfully")

}

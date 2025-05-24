package handlers

import (
	"back/config"
	"back/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddCarHandler(w http.ResponseWriter, r *http.Request){
	var car models.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	id, err := config.RedisDataDb.Incr(r.Context(), "car:id:counter").Result()
	car.ID = fmt.Sprintf("%d", id)

	key := "car:" + car.ID
	err = config.RedisDataDb.HSet(r.Context(), key, map[string]interface{}{
		"brand": car.Brand,
		"model": car.Model,
	}).Err()
	if err != nil {
		http.Error(w, "Failed to save car", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteCarHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}

	key := "car:" + id
	err := config.RedisDataDb.Del(r.Context(), key).Err()
	if err != nil {
		http.Error(w, "Failed to delete car", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

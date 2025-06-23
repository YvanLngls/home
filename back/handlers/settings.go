package handlers

import (
	"back/config"
	"back/utils"
	"context"
	"encoding/json"
	"net/http"
)

func GetUserInfos(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	claims, err := utils.ValidateJWT(cookie.Value)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	user := claims.Username
	ctx := context.Background()
	
	// Récupère l’ID via username
	userId, err := config.RedisUserDb.Get(ctx, "username:"+user).Result()
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Récupère la date de création
	createdAt, err := config.RedisDataDb.HGet(ctx, "user:"+userId, "creation").Result()
	if err != nil {
  	http.Error(w, "Failed to get creation date", http.StatusInternalServerError)
		return
	}

	// Récupère la date de dernière connection
	lastConnection, err := config.RedisDataDb.HGet(ctx, "user:"+userId, "lastConnection").Result()
	if err != nil {
  	http.Error(w, "Failed to get last connexion date", http.StatusInternalServerError)
		return
	}

	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
	// json.NewEncoder(w).Encode(map[string]string{
		"username":    user,
		"createdAt":  createdAt,
		"lastConnection":  lastConnection,
	})
}

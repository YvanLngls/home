package handlers

import (
	"back/config"
	"back/models"
	"back/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	// Vérifie si username existe déjà
	_, err := config.RedisUserDb.Get(ctx, "username:"+user.Username).Result()
	if err == nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	// Génère un ID unique
	id, err := config.RedisUserDb.Incr(ctx, "user:next_id").Result()
	if err != nil {
		http.Error(w, "Failed to generate user ID", http.StatusInternalServerError)
		return
	}
	userId := fmt.Sprintf("%d", id)

	// Hash le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Stocke l'utilisateur dans un hash
	userKey := "user:" + userId
	err = config.RedisUserDb.Set(ctx, userKey, string(hashedPassword), 0).Err()
	if err != nil {
		http.Error(w, "Failed to store user", http.StatusInternalServerError)
		return
	}

	// Stocke l'utilisateur
	err = config.RedisDataDb.HSet(ctx, userKey, map[string]interface{}{
		"username":   user.Username,
		"creation": time.Now().Format(time.RFC3339),
		"lastConnection": time.Now().Format(time.RFC3339),
		"joinedSimulations": "[]",
	}).Err()
	if err != nil {
		http.Error(w, "Failed to store user", http.StatusInternalServerError)
		return
	}

	// Crée l’index username → userId
	err = config.RedisUserDb.Set(ctx, "username:"+user.Username, userId, 0).Err()
	if err != nil {
		http.Error(w, "Failed to index username", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User registered successfully")
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	// Trouve l’ID via le username
	userId, err := config.RedisUserDb.Get(ctx, "username:"+user.Username).Result()
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Récupère les données utilisateur
	userKey := "user:" + userId
	storedPassword, err := config.RedisUserDb.Get(ctx, userKey).Result()
	if err != nil {
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	tokenString, expirationTime, err := utils.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	err = config.RedisDataDb.HSet(ctx, userKey, map[string]interface{}{
		"lastConnection": time.Now().Format(time.RFC3339),
	}).Err()
	if err != nil {
		http.Error(w, "Failed to store last connection date", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    user.Username,
		Expires:  expirationTime,
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logged in successfully",
		"token":   tokenString,
	})
}


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

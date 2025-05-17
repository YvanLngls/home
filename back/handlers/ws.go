package handlers

import (
	"net/http"
	"github.com/gorilla/websocket"
	"back/utils"
	"log"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // adjust for security
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Get token from query or cookie
	tokenStr, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
		return
	}

	// 2. Validate JWT token
	claims, err := utils.ValidateJWT(tokenStr.Value)
	if err != nil {
		http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
		return
	}

	// 3. Upgrade HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	log.Println("WebSocket connected:", claims.Username)

	// 4. Example: read messages or send data
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received from %s: %s", claims.Username, string(msg))

		// Echo message back
		err = conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

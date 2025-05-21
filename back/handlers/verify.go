package handlers

import (
	"encoding/json"
	"net/http"
)

func VerifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"valid": true})
}

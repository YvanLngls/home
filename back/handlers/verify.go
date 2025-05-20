package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"
)

func VerifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Token received")
	json.NewEncoder(w).Encode(map[string]bool{"valid": true})
}

package controller

import (
	"encoding/json"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("pong")
}

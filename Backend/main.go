package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	// Allow React frontend to connect
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

	w.Header().Set("Content-Type", "application/json")

	res := Response{
		Message: "Hello from Go backend",
	}

	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/api", handler)

	http.ListenAndServe(":8080", nil)
}

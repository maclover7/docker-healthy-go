package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIResponse struct {
	Services map[string]string `json:"services"`
}

func main() {
	fmt.Printf("Starting server...\n")

	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":9292", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	services := &APIResponse{
		Services: map[string]string{"macluster": "Up 23 seconds (healthy)"}}
	serializedServices, _ := json.Marshal(services)

	fmt.Fprintf(w, string(serializedServices))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

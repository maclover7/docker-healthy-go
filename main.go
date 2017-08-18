package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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

// HTTP handlers

func apiHandler(w http.ResponseWriter, r *http.Request) {
	services := &APIResponse{
		Services: ContainerList()}
	serializedServices, _ := json.Marshal(services)

	fmt.Fprintf(w, string(serializedServices))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

// Helper functions

func ContainerList() map[string]string {
	var services map[string]string
	services = make(map[string]string)

	cli, err := client.NewEnvClient()
	if err != nil {
		return services
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return services
	}

	for _, container := range containers {
		services[container.Image] = container.Status
	}

	return services
}

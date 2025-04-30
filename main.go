package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// getResourcesHandler handles GET requests and returns a list of resources
func getResourcesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Resources); err != nil {
		http.Error(w, "Failed to encode resources", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/api/resources", getResourcesHandler)
	log.Println("Server starting at :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

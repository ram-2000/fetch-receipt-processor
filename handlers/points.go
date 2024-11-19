package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
)

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the receipt ID from the URL path
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/receipts/"), "/")
	if len(pathParts) < 1 || pathParts[0] == "" {
		http.Error(w, "Invalid receipt ID", http.StatusBadRequest)
		return
	}

	id := pathParts[0] // Extract the receipt ID without /points suffix
	log.Printf("Fetching receipt with ID: %s", id)

	// Retrieve data from Memcached
	item, err := MemClient.Get(id)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			log.Printf("Receipt with ID %s not found in Memcached", id)
			http.Error(w, "Receipt not found", http.StatusNotFound)
		} else {
			log.Printf("Error retrieving receipt from Memcached: %v", err)
			http.Error(w, "Failed to retrieve receipt", http.StatusInternalServerError)
		}
		return
	}

	log.Printf("Successfully retrieved data for ID: %s", id)

	// Deserialize JSON data
	var data map[string]interface{}
	if err := json.Unmarshal(item.Value, &data); err != nil {
		log.Printf("Failed to decode receipt data: %v", err)
		http.Error(w, "Failed to decode receipt data", http.StatusInternalServerError)
		return
	}

	// Extract precomputed points
	points := int(data["points"].(float64))
	log.Printf("Points for receipt ID %s: %d", id, points)

	// Return the points
	response := map[string]int{"points": points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

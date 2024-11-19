package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"receipt-service/models"
	"receipt-service/utils"
	"github.com/bradfitz/gomemcache/memcache"
)
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
    var receipt models.Receipt

    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    id := uuid.New().String()
    points := utils.CalculatePoints(&receipt)

    // Serialize data to JSON
    receiptData := map[string]interface{}{
        "receipt": receipt,
        "points":  points,
    }
    jsonData, _ := json.Marshal(receiptData)

    // Store data in Memcached
    err := MemClient.Set(&memcache.Item{
        Key:   id,
        Value: jsonData,
    })
    if err != nil {
        log.Printf("Failed to store receipt in Memcached: %v", err)
        http.Error(w, "Failed to store receipt", http.StatusInternalServerError)
        return
    }

    log.Printf("Successfully stored receipt with ID: %s", id)

    response := map[string]string{"id": id}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"go-api-test/database"
	"go-api-test/models"
)

// getItems handles GET requests to /items
func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []models.Item
	database.DB.Find(&items)

	json.NewEncoder(w).Encode(items)
}

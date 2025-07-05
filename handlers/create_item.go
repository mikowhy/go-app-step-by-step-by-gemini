package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"go-api-test/database"
	"go-api-test/models"
)

// createItem handles POST requests to /items
func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item.ID = uuid.New().String()

	if result := database.DB.Create(&item); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

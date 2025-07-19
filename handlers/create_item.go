package handlers

import (
	"bytes"
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

	// Use a buffer to handle potential empty body issues
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	if buf.Len() == 0 {
		http.Error(w, `{"error": "Request body cannot be empty"}`, http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(buf.Bytes(), &item); err != nil {
		http.Error(w, `{"error": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}

	// --- Validation ---
	if item.Name == "" {
		http.Error(w, `{"error": "Item name cannot be empty"}`, http.StatusBadRequest)
		return
	}
	if item.Description == "" {
		http.Error(w, `{"error": "Item description cannot be empty"}`, http.StatusBadRequest)
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

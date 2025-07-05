package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go-api-test/database"
	"go-api-test/models"
)

// updateItem handles PUT requests to /items/{id}
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedItem models.Item
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var item models.Item
	if result := database.DB.First(&item, "ID = ?", id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	// Update fields
	item.Name = updatedItem.Name
	item.Description = updatedItem.Description

	if result := database.DB.Save(&item); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

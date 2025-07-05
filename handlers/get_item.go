package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go-api-test/database"
	"go-api-test/models"
)

// getItem handles GET requests to /items/{id}
func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var item models.Item
	if result := database.DB.First(&item, "ID = ?", id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

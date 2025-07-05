package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-api-test/database"
	"go-api-test/models"
)

// deleteItem handles DELETE requests to /items/{id}
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var item models.Item
	if result := database.DB.First(&item, "ID = ?", id); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	if result := database.DB.Delete(&item); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

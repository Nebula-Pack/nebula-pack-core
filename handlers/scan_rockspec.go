package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nebula-Pack/nebula-pack-core/models"
)

func HandleScanRockspec(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Path string `json:"path"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Error parsing JSON body", http.StatusBadRequest)
		return
	}

	dependencies, err := models.ScanRockspec(request.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error scanning .rockspec file: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dependencies)
}

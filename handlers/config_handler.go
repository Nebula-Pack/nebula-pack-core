package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nebula-Pack/nebula-pack-core/models"
)

func HandleNebulaConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the body of the request
	var config models.NebulaConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, "Error parsing JSON body", http.StatusBadRequest)
		return
	}

	// Print out the parsed config (for testing)
	fmt.Printf("Received Nebula Config:\n%+v\n", config)

	// Process the Nebula config (for example, deduct dependencies)
	err := config.Process()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing Nebula Config: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Nebula Config received and processed successfully\n")
}

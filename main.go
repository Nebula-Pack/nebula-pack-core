package main

import (
	"log"
	"net/http"

	"github.com/Nebula-Pack/nebula-pack-core/handlers"
)

func main() {
	// Registering handlers
	http.HandleFunc("/nebula-config", handlers.HandleNebulaConfig)

	// Starting server
	log.Fatal(http.ListenAndServe(":7777", nil))
}

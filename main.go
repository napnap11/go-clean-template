package main

import (
	"log"

	"github.com/napnap11/go-clean-template/internal/app/route"
	"github.com/napnap11/go-clean-template/internal/pkg/config"
	"github.com/napnap11/go-clean-template/internal/pkg/util/https"
)

func main() {
	// Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Initialize router
	router := route.NewRouter()

	// Create and run server
	addr := ":" + cfg.HTTP.Port
	server := https.NewServer(addr, router)
	if err := server.Run(); err != nil {
		log.Fatalf("Server error: %s", err)
	}
}

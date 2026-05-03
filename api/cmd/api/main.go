package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/RezaBagusS/loka-app/api/internal/config"
	"github.com/RezaBagusS/loka-app/api/pkg/database"
)

func main() {
	// 1. Load Configuration (Viper)
	cfg, err := config.LoadConfig("../../")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// 2. Initialize Database (sqlx)
	db, err := database.InitDB(cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	if err != nil {
		log.Printf("Warning: Database connection failed: %v", err)
	} else {
		defer db.Close()
	}

	// 3. Initialize Chi Router
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		res := map[string]string{
			"message": "pong",
			"app":     "LokaApp API (Chi + Sqlx + Viper)",
			"status":  "active",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	// Server Port
	port := cfg.ServerPort
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":" + port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

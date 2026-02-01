package main

import (
	"log"
	"net/http"
	"kanban-chik/internal/config"
	"kanban-chik/internal/db"
	"kanban-chik/internal/handlers"
)

func main() {
	log.Println("starting app...")

	cfg := config.Load()
	database := db.Connect(cfg.DatabaseURL)
	defer database.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.Health)

	log.Println("http server started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

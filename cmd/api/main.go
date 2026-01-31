package main

import (
	"log"
	"net/http"

	"ChaosApi/internal/chaos"
	"ChaosApi/internal/handlers"
	"ChaosApi/internal/middleware"
	"ChaosApi/internal/server"
)

func main() {
	store := chaos.NewInMemoryStore()
	engine := chaos.NewEngine(store)
	chaosMiddleware := middleware.Chaos(engine)

	r := server.NewRouter(chaosMiddleware)

	// Endpoints
	r.Get("/users", handlers.UsersHandler)
	r.Mount("/chaos", handlers.ChaosRoutes(store)) // CRUD de regras

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

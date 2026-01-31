package handlers

import (
	"encoding/json"
	"net/http"

	"ChaosApi/internal/chaos"

	"github.com/go-chi/chi/v5"
)

func ChaosRoutes(store chaos.Store) http.Handler {
	r := chi.NewRouter()

	// GET todos
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		cfgs := store.List()
		json.NewEncoder(w).Encode(cfgs)
	})

	// CREATE
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var cfg chaos.Config
		if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id := store.Create(&cfg)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"id": id})
	})

	// DELETE
	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if deleted := store.Delete(id); !deleted {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	return r
}

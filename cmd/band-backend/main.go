package main

import (
	"net/http"

	"github.com/arturbaccarin/band-backend/config"
	"github.com/arturbaccarin/band-backend/internal/infra/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	cfg := config.LoadConfig()

	db := database.OpenConnection(
		cfg.DBHost,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBName,
		cfg.WebServerPort,
	)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}

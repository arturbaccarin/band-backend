package main

import (
	"net/http"

	"github.com/arturbaccarin/band-backend/config"
	"github.com/arturbaccarin/band-backend/internal/infra/database"
	"github.com/arturbaccarin/band-backend/webserver/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	cfg := config.LoadConfig()

	db := database.OpenConnection(
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBName,
		cfg.DBPort,
	)

	bandDB := database.NewBand(db)
	bandHandler := handler.NewBandHandler(bandDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/bands", func(r chi.Router) {
		r.Post("/", bandHandler.Create)
		r.Get("/{ID}", bandHandler.GetByID)
	})

	http.ListenAndServe(":8000", r)
}

package main

import (
	"net/http"

	"github.com/arturbaccarin/band-backend/config"
	_ "github.com/arturbaccarin/band-backend/docs"
	"github.com/arturbaccarin/band-backend/internal/infra/database"
	"github.com/arturbaccarin/band-backend/internal/infra/webserver"
	"github.com/arturbaccarin/band-backend/internal/infra/webserver/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			API Band
//	@version		1.0
//	@description	Band API for my personal project
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Artur Baccarin

// @host						localhost:8000
// @BasePath					/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
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

	userDB := database.NewUser(db)
	userHandler := handler.NewUserHander(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			panic(err)
		}
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	r.Route("/bands", func(r chi.Router) {
		r.Use(webserver.JWTAuthenticator)
		r.Post("/", bandHandler.Create)
		r.Get("/{ID}", bandHandler.GetByID)
		r.Delete("/{ID}", bandHandler.DeleteByID)
		r.Put("/{ID}", bandHandler.UpdateByID)
		r.Get("/", bandHandler.GetList)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.Create)
		r.Post("/signin", userHandler.SignIn)
	})

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
	}
}

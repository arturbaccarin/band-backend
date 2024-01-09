package webserver

import (
	"net/http"

	"github.com/arturbaccarin/band-backend/pkg/auth"
)

func JWTAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		err := auth.ValidateJWT(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			_, err := w.Write([]byte(err.Error()))
			if err != nil {
				panic(err)
			}

			return
		}

		next.ServeHTTP(w, r)
	})
}

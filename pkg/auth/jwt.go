package auth

import (
	"time"

	"github.com/go-chi/jwtauth"
)

var cfg *AuthConfig

func init() {
	cfg = LoadConfig()
}

func GenerateJWT(sub string) (string, error) {
	secretKey := []byte(cfg.JWTSecretKey)
	tokenAuth := jwtauth.New("HS256", []byte(secretKey), nil)

	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{
		"sub": sub,
		"exp": time.Now().Add(time.Duration(cfg.JWTExpiresIn) * time.Second).Unix(),
	})

	return tokenString, nil
}

package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var cfg *AuthConfig

func init() {
	cfg = LoadConfig()
}

func GenerateJWT(sub string) (string, error) {
	secretKey := []byte(cfg.JWTSecretKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": sub,
			"exp": time.Now().Add(time.Duration(cfg.JWTExpiresIn) * time.Second).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var cfg *AuthConfig
var secretKey []byte

func init() {
	cfg = LoadConfig()
	secretKey = []byte(cfg.JWTSecretKey)
}

func GenerateJWT(sub string) (string, error) {
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
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

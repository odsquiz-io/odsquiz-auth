// internal/utils/Token.go: setup the token feature to allow routing through protected routes
package utils

import (
	"time"
	"log"
	
	"github.com/kauanpecanha/odsquiz-auth/pkg/config"
	
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID string, email string) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	secretKey := []byte(cfg.JWTSecret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{ 
			"sub": userID,
			"email": email,
			"exp": time.Now().Add(24 * time.Hour).Unix(),
			"iat": time.Now().Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
	return "", err
	}

 return tokenString, nil
}
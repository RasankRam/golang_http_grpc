package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key") // Change this to a secure secret in production

// JWT claims struct
type Claims struct {
	Login string `json:"login"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAuthToken(login string, role string) (string, error) {
	//expiration := time.Now().Add(expirationTime) // если в будущем понадобиться связка refresh-access token
	claims := &Claims{
		Login: login,
		Role:  role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

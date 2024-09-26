package utils

import (
	"github.com/mdobak/go-xerrors"
	"strings"
)

func BearerToken(authHeader string) (string, error) {
	token, err := parseToken(authHeader)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Проверка токенов на authTokens
func parseToken(rawToken string) (string, error) {
	if strings.HasPrefix(rawToken, "Bearer ") {
		token := strings.TrimPrefix(rawToken, "Bearer ")

		return token, nil
	}
	return "", xerrors.New("invalid Token")
}

package mw

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"net/http"
	"strings"
	"todo-list/internal/app/repository"
	logger2 "todo-list/internal/logger"
	"todo-list/internal/utils"
)

// Middleware for authorization
func AuthMiddleware(next http.Handler, db *sqlx.DB, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.BearerToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		repo := repository.NewTokenRepo(db)

		tokenEntity, err := repo.TokenByToken(token)

		fmt.Println("tokenEntity", tokenEntity)
		fmt.Println("tokenEntity_err", err)

		if err != nil || tokenEntity == nil {
			utils.LogErrorContext(logger, r.Context(), err)
			http.Error(w, "failed to find token", http.StatusUnauthorized)
			return
		}

		ctx := logger2.AppendCtx(r.Context(), slog.Int(utils.ContextUserKey, tokenEntity.UserId))

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// Проверка токенов на authTokens
func parseToken(rawToken string) (string, error) {
	if strings.HasPrefix(rawToken, "Bearer ") {
		token := strings.TrimPrefix(rawToken, "Bearer ")

		return token, nil
	}
	return "", errors.New("invalid Token")
}

package routes

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
	"net/http"
	"todo-list/internal/app/controllers"
	"todo-list/internal/app/repository"
	"todo-list/internal/app/services"
)

func SetupAuthRoutes(router *http.ServeMux, db *sqlx.DB, logger *slog.Logger) {
	authRouter := http.NewServeMux()
	userRepo := repository.NewUserRepo(db)
	tokenRepo := repository.NewTokenRepo(db)
	roleRepo := repository.NewRoleRepo(db)
	service := services.NewAuthService(userRepo, roleRepo, tokenRepo, logger)
	controller := controllers.NewAuthController(service)

	authRouter.HandleFunc("POST /register", controller.Register)
	authRouter.HandleFunc("POST /login", controller.Login)
	authRouter.HandleFunc("POST /logout", controller.Logout)

	router.Handle("/auth/", http.StripPrefix("/auth", authRouter))
}

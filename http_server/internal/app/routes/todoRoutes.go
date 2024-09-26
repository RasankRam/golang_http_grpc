package routes

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
	"net/http"
	"todo-list/internal/app/controllers"
	"todo-list/internal/app/mw"
	"todo-list/internal/app/repository"
	"todo-list/internal/app/services"
)

func SetupTodoRoutes(router *http.ServeMux, db *sqlx.DB, logger *slog.Logger) {
	repo := repository.NewTodoRepo(db)
	service := services.NewTodoService(repo, logger)
	controller := controllers.NewTodoController(service, logger)

	todoRouter := http.NewServeMux()

	todoRouter.HandleFunc("GET /", controller.Todos)
	todoRouter.HandleFunc("GET /{id}", controller.Todo)
	todoRouter.HandleFunc("POST /", controller.AddTodo)
	todoRouter.HandleFunc("DELETE /{id}", controller.DeleteTodo)
	todoRouter.HandleFunc("PUT /{id}", controller.UpdateTodo)

	todoRouterMw := mw.AuthMiddleware(todoRouter, db, logger)

	router.Handle("/todos/", http.StripPrefix("/todos", todoRouterMw))
}

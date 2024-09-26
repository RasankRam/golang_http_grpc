package controllers

import (
	"encoding/json"
	"github.com/mdobak/go-xerrors"
	"log/slog"
	"net/http"
	"strconv"
	todo_requests "todo-list/internal/app/requests/todo"
	"todo-list/internal/app/services"
	"todo-list/internal/utils"
)

type TodoController struct {
	service *services.TodoService
	logger  *slog.Logger
}

func NewTodoController(service *services.TodoService, logger *slog.Logger) *TodoController {
	return &TodoController{
		service: service,
		logger:  logger,
	}
}

func (c *TodoController) Todo(w http.ResponseWriter, r *http.Request) {
	todoIdStr := r.PathValue("id")

	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil || todoId == 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := c.service.Todo(r.Context(), todoId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Handle error
		return
	}
}

func (c *TodoController) Todos(w http.ResponseWriter, r *http.Request) {

	todos, err := c.service.Todos(r.Context())

	if err != nil {
		http.Error(w, "Failed to receive todos", http.StatusInternalServerError)
		return
	}

	// Convert the slice of structs to JSON
	w.Header().Set("Content-Type", "application/json") // Set content type to JSON
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Handle error
		return
	}
}

func (c *TodoController) AddTodo(w http.ResponseWriter, r *http.Request) {

	addTodoReq, err := todo_requests.CreateAddTodoReq(r.Body, &r.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value(utils.ContextUserKey).(slog.Attr)
	if !ok {
		utils.LogErrorContext(c.logger, r.Context(), xerrors.New("user ID not found"))
		http.Error(w, "failed to find auth user", http.StatusUnauthorized)
	}
	userIdVal := int(userId.Value.Int64())

	todoId, err := c.service.AddTodo(r.Context(), addTodoReq.Title, addTodoReq.Dsc, userIdVal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(todoId)))
}

func (c *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoIdStr := r.PathValue("id")
	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil || todoId == 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := c.service.DeleteTodo(r.Context(), todoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Handle error
		return
	}
}

func (c *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoIdStr := r.PathValue("id")
	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil || todoId == 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	chgTodoReq, err := todo_requests.CreateChgTodoReq(r.Body, &r.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value(utils.ContextUserKey).(slog.Attr)
	if !ok {
		utils.LogErrorContext(c.logger, r.Context(), xerrors.New("user ID not found"))
		http.Error(w, "failed to find auth user", http.StatusUnauthorized)
	}
	userIdVal := int(userId.Value.Int64())

	modifiedTodo, err := c.service.UpdateTodo(r.Context(), todoId, chgTodoReq.Title, chgTodoReq.Dsc, userIdVal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(modifiedTodo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Handle error
		return
	}
}

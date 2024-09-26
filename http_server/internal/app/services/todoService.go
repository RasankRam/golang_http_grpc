package services

import (
	"context"
	"errors"
	pb "example.com/task_platform_proto/gen_go"
	"fmt"
	"github.com/mdobak/go-xerrors"
	"log/slog"
	"todo-list/internal/app/repository"
	"todo-list/internal/models"
	"todo-list/internal/utils"
)

type TodoService struct {
	repo   *repository.TodoRepo
	logger *slog.Logger
}

func NewTodoService(repo *repository.TodoRepo, logger *slog.Logger) *TodoService {
	return &TodoService{
		repo:   repo,
		logger: logger,
	}
}

func (t *TodoService) Todos(ctx context.Context) ([]models.Todo, error) {
	todos, err := t.repo.Todos()

	if err != nil {
		utils.LogErrorContext(t.logger, ctx, err)
		return nil, errors.New("failed to fetch todos")
	}

	return todos, err
}

func (t *TodoService) Todo(ctx context.Context, id int) (*models.Todo, error) {
	todo, err := t.repo.Todo(id)

	if err != nil {
		utils.LogErrorContext(t.logger, ctx, err)
		return nil, errors.New("failed to fetch todo")
	}

	return todo, err
}

func (t *TodoService) AddTodo(ctx context.Context, title string, dsc string, authUserId int) (int, error) {
	todoId, err := t.repo.AddTodo(title, dsc, authUserId)

	if err != nil {
		utils.LogErrorContext(t.logger, ctx, err)
		return 0, errors.New("failed to add todo")
	}

	grpcClient, ok := utils.GetGrpcClientFromContext(ctx)
	if !ok {
		err = xerrors.New("gRPC client not found in context")
		utils.LogErrorContext(t.logger, ctx, err)
		return 0, errors.New(err.Error())
	}
	fmt.Println("grpcClient", grpcClient)
	go func() {
		req := &pb.TodoRequest{
			Name:  title,
			Price: 100.00,
		}

		fmt.Println("req", req)

		// Call the gRPC service
		res, err := grpcClient.ProcessTodo(context.Background(), req)
		if err != nil {
			fmt.Println("grpcClient_ProcessTodo_err", err)
			return
		}
		fmt.Println("discounted_price", res.GetDiscountedPrice())
	}()

	return todoId, err
}

func (t *TodoService) DeleteTodo(ctx context.Context, id int) (*models.Todo, error) {
	todo, err := t.repo.DeleteTodo(id)

	if err != nil {
		utils.LogErrorContext(t.logger, ctx, err)
		return nil, errors.New("failed to delete todo")
	}

	return todo, err
}

func (t *TodoService) UpdateTodo(ctx context.Context, todoId int, title string, dsc string, authUserId int) (*models.Todo, error) {
	todo, err := t.repo.UpdateTodo(todoId, title, dsc, authUserId)

	if err != nil {
		utils.LogErrorContext(t.logger, ctx, err)
		return nil, errors.New("failed to update todo")
	}

	return todo, err
}

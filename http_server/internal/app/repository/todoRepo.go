package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-list/internal/models"
)

type TodoRepo struct {
	db *sqlx.DB
}

func NewTodoRepo(db *sqlx.DB) *TodoRepo {
	return &TodoRepo{db: db}
}

func (t *TodoRepo) Todos() ([]models.Todo, error) {
	todos := []models.Todo{}

	err := t.db.Select(&todos, "SELECT * FROM todos")
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *TodoRepo) Todo(id int) (*models.Todo, error) {
	todo := models.Todo{}
	err := t.db.Get(&todo, "SELECT * FROM todos WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *TodoRepo) UpdateTodo(todoId int, title string, dsc string, createdBy int) (*models.Todo, error) {

	todo := models.Todo{}
	err := t.db.Get(&todo, "update todos set title=$1,dsc=$2,updated_by=$3,updated_at=now() where id=$4 returning *", title, dsc, createdBy, todoId)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *TodoRepo) DeleteTodo(id int) (*models.Todo, error) {
	todo := models.Todo{}
	err := t.db.Get(&todo, "delete from todos where id=$1 returning *", id)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (t *TodoRepo) AddTodo(title string, dsc string, createdBy int) (int, error) {

	query := `insert into todos(title,dsc,created_by,created_at) 
				values (:title, :dsc, :created_by, now()) returning id`

	stmt, err := t.db.PrepareNamed(query)
	if err != nil {
		return 0, err
	}

	var todoId int
	err = stmt.Get(&todoId, map[string]interface{}{
		"title":      title,
		"dsc":        dsc,
		"created_by": createdBy,
	})

	if err != nil {
		return 0, err
	}

	return todoId, nil
}

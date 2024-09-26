package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-list/internal/models"
)

type RoleRepo struct {
	db *sqlx.DB
}

func NewRoleRepo(db *sqlx.DB) *RoleRepo {
	return &RoleRepo{
		db: db,
	}
}

func (r *RoleRepo) RoleByNm(nm string) (*models.Role, error) {
	role := models.Role{}
	err := r.db.Get(&role, "SELECT * FROM roles where nm = $1", nm)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

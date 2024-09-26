package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-list/internal/models"
)

type TokenRepo struct {
	db *sqlx.DB
}

func NewTokenRepo(db *sqlx.DB) *TokenRepo {
	return &TokenRepo{
		db: db,
	}
}

func (r *TokenRepo) TokenByToken(token string) (*models.Token, error) {
	var tokenEntity models.Token
	err := r.db.Get(&tokenEntity, "select * from tokens where token = $1", token)
	if err != nil {
		return nil, err
	}

	return &tokenEntity, nil
}

func (r *TokenRepo) DeleteToken(token string) (string, error) {
	// Delete the refresh token from the database
	var deletedToken string
	err := r.db.Get(&deletedToken, "delete from tokens where token = $1 returning token", token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *TokenRepo) CreateToken(token string, ip string, user_id int, user_agent string) (int, error) {
	var tokenId int
	err := r.db.Get(&tokenId, "insert into tokens(token,ip,user_id,user_agent) values ($1,$2,$3,$4) returning id", token, ip, user_id, user_agent)
	if err != nil {
		return 0, err
	}
	return tokenId, nil
}

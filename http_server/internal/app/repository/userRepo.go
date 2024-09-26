package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mdobak/go-xerrors"
	"todo-list/internal/models"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) UserByLogin(login string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user,
		`select u.id, u.login, u.password, r.nm as role 
    		   from users u
			   join roles r on r.id = u.role_id 
    		   where u.login = $1
	`, login)
	fmt.Println("userByLogin", user)
	fmt.Println("userByLogin_err", err)
	if err != nil {
		return nil, xerrors.New("Invalid to get user by login")
	}
	return &user, nil
}

func (r *UserRepo) CreateUserRegister(login string, hashedPass string, role string) (int, error) {
	var userId int
	err := r.db.Get(&userId,
		"insert into users(login, password, role_id) values ($1, $2, (select id from roles where nm = $3)) returning id",
		login, hashedPass, role)

	if err != nil {
		return 0, err
	}
	return userId, nil
}

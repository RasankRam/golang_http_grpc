package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"todo-list/internal/utils"
)

func OpenSqlxViaPgxConnPool() *sqlx.DB {
	dbName := utils.MustGetenv("DB_NAME")
	dbUser := utils.MustGetenv("DB_USER")
	dbPass := utils.MustGetenv("DB_PASS")
	var dbHost string
	var dbPort string
	inDocker := os.Getenv("IN_DOCKER")
	if inDocker != "" {
		dbHost = "db"
		dbPort = "5432"
	} else {
		dbHost = utils.MustGetenv("DB_HOST")
		dbPort = utils.MustGetenv("DB_PORT")
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	dbPool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}

	dbConn := stdlib.OpenDBFromPool(dbPool)

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return sqlx.NewDb(dbConn, "pgx")
}

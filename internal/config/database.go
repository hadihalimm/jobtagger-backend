package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type Database interface {
}

type database struct {
	db *pgx.Conn
}

var dbInstance *database

func ConnectToDatabase() Database {
	if dbInstance != nil {
		return dbInstance
	}

	connStr := "postgres://postgres:postgres@localhost:5432/jobtagger"
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &database{db: db}

	return dbInstance
}

package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Database interface {
	CreateTables()
}

type database struct {
	db *pgx.Conn
}

var dbInstance *database

func ConnectToDatabase() Database {
	if dbInstance != nil {
		return dbInstance
	}

	connStr := "postgres://postgres:postgres@localhost:5432/job-tagger"
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &database{db: db}

	return dbInstance
}

func (d *database) CreateTables() {
	workDir, _ := os.Getwd()
	sql, readError := os.ReadFile(fmt.Sprintf("%s/internal/config/schema.sql", workDir))
	if readError != nil {
		panic(fmt.Sprintln(readError))
	}

	_, execError := d.db.Exec(context.Background(), string(sql))
	if execError != nil {
		panic(fmt.Sprintln(execError))
	}
}

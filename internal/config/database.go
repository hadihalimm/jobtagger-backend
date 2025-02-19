package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	Pgx *pgx.Conn
}

var dbInstance *Database

func ConnectToDatabase() *Database {
	if dbInstance != nil {
		return dbInstance
	}

	connStr := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &Database{Pgx: conn}

	return dbInstance
}
func (database *Database) Close() {
	database.Pgx.Close(context.Background())
}

func (database *Database) CreateAllTables() {
	workDir, _ := os.Getwd()
	sql, readError := os.ReadFile(fmt.Sprintf("%s/internal/config/schema.sql", workDir))
	if readError != nil {
		panic(fmt.Sprintln(readError))
	}

	_, execError := database.Pgx.Exec(context.Background(), string(sql))
	if execError != nil {
		panic(fmt.Sprintln(execError))
	}
}

func (database *Database) DropAllTables() {
	_, execError := database.Pgx.Exec(context.Background(), "DROP SCHEMA IF EXISTS public CASCADE; CREATE SCHEMA public;")
	if execError != nil {
		panic(fmt.Sprintln(execError))
	}
}

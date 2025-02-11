package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/hadihalimm/jobtagger-backend/internal/config"
)

type Server struct {
	port int
	db   config.Database
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := config.ConnectToDatabase()
	db.CreateTables()

	NewServer := &Server{
		port: port,
		db:   db,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server
}

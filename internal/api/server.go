package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/hadihalimm/jobtagger-backend/internal/config"
	"github.com/hadihalimm/jobtagger-backend/internal/handler"
	"github.com/hadihalimm/jobtagger-backend/internal/repo"
	"github.com/hadihalimm/jobtagger-backend/internal/service"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port                  int
	db                    *config.Database
	authHandler           *handler.AuthHandler
	jobApplicationHandler *handler.JobApplicationHandler
	interviewHandler      *handler.InterviewHandler
	contactHandler        *handler.ContactHandler
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := config.ConnectToDatabase()
	// db.DropAllTables()
	// db.CreateAllTables()

	service.InitAuth()
	userRepo := repo.NewUserRepo(db)
	refreshTokenRepo := repo.NewRefreshTokenRepo(db)
	jobApplicationRepo := repo.NewJobApplicationRepo(db)
	interviewRepo := repo.NewInterviewRepo(db)
	contactRepo := repo.NewContactRepo(db)

	authService := service.NewAuthService(userRepo, refreshTokenRepo)
	jobApplicationService := service.NewJobApplicationService(jobApplicationRepo)
	interviewService := service.NewInterviewService(interviewRepo)
	contactService := service.NewContactService(contactRepo)

	authHandler := handler.NewAuthHandler(authService)
	jobApplicationHandler := handler.NewJobApplicationHandler(jobApplicationService)
	interviewHandler := handler.NewInterviewHandler(interviewService)
	contactHandler := handler.NewContactHandler(contactService)

	NewServer := &Server{
		port:                  port,
		db:                    db,
		authHandler:           authHandler,
		jobApplicationHandler: jobApplicationHandler,
		interviewHandler:      interviewHandler,
		contactHandler:        contactHandler,
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

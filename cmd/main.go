package main

import (
	"context"
	"github.com/ashurov-imomali/qa-api/internal/app"
	"github.com/ashurov-imomali/qa-api/internal/db"
	"github.com/ashurov-imomali/qa-api/internal/handlers"
	"github.com/ashurov-imomali/qa-api/internal/repository"
	"github.com/ashurov-imomali/qa-api/internal/service"
	"github.com/ashurov-imomali/qa-api/pkg/logger"
	"os"
	"os/signal"
	"time"
)

func main() {
	log := logger.New()

	// Get DB DSN
	dbDSN := os.Getenv("DB_DSN")
	if dbDSN == "" {
		dbDSN = "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable"
	}

	// Get APP port
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Run migrations
	if err := db.RunMigrations(dbDSN); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Connect to DB
	pg, err := db.NewConnection(dbDSN)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize repository
	repo := repository.NewRepository(pg)

	// Initialize services
	answerService := service.NewAnswerService(repo, log)
	questionService := service.NewQuestionService(repo, log)

	// Initialize HTTP handlers
	handler := handlers.New(questionService, answerService, log)

	// Create server
	server := app.NewServer(":"+port, handler)

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Infof("Server starting on port %s", port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	<-stop
	log.Infof("%s", "Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Errorf("Graceful shutdown failed: %v", err)
	} else {
		log.Infof("%s", "Server stopped gracefully")
	}
}

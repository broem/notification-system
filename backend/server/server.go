package server

import (
	"flag"
	"log"

	"github.com/broem/notification-system/backend/internal/config"
	"github.com/broem/notification-system/backend/internal/handler"
	"github.com/broem/notification-system/backend/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(cfg *config.Config) {
	migrateOnly := flag.Bool("migrate-only", false, "Apply migrations and exit")
	flag.Parse()

	// Set up database connection
	db, err := repository.SetupDatabase(cfg.Database)
	if err != nil {
		panic(err)
	}

	// Apply migrations
	migrationsPath := "./migrations"
	err = repository.ApplyMigrations(db, migrationsPath)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	if *migrateOnly {
		return
	}

	// Create repositories
	notificationRepo := repository.CreateNotificationRepository(db)

	// Create handlers
	notificationHandler := handler.NewNotificationHandler(notificationRepo)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Group routes
	apiv1 := e.Group("/api/v1")

	// Routes
	apiv1.GET("/notifications", notificationHandler.GetAllNotifications)

	// Start server
	e.Logger.Fatal(e.Start(cfg.Server.Address))
}
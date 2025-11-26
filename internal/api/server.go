package api

import (
	"go-microservices/config"
	"go-microservices/internal/api/rest"
	"go-microservices/internal/api/rest/handlers"
	"go-microservices/internal/domain"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("database connection error %v\n", err)
	}

	log.Println("database connected")

	// Run migration
	db.AutoMigrate(&domain.User{})

	log.Print(db)

	rh := &rest.RestHandler{
		App: app,
		DB:  db,
	}

	setupRoutes(rh)
	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	//user handler
	handlers.SetupUserRoutes(rh)

	//other
}

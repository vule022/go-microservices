package api

import (
	"go-ecommerce/config"
	"go-ecommerce/internal/api/rest"
	"go-ecommerce/internal/api/rest/handlers"
	"go-ecommerce/internal/domain"
	"go-ecommerce/internal/helper"
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
	err = db.AutoMigrate(&domain.User{}, &domain.BankAccount{})

	if err != nil {
		log.Fatalf("migration error %v", err.Error())
	}

	log.Println("migration successful")

	auth := helper.SetupAuth(config.AppSecret)

	log.Print(db)

	rh := &rest.RestHandler{
		App:    app,
		DB:     db,
		Auth:   auth,
		Config: config,
	}

	setupRoutes(rh)
	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	//user handler
	handlers.SetupUserRoutes(rh)

	//other
}

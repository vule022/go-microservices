package main

import (
	"go-microservices/config"
	"go-microservices/internal/api"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)
}

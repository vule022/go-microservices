package service

import (
	"go-ecommerce/config"
	"go-ecommerce/internal/helper"
	"go-ecommerce/internal/repository"
)

type CatalogService struct {
	Repo   repository.GatalogRepository
	Auth   helper.Auth
	Config config.AppConfig
}

package repository

import (
	"gorm.io/gorm"
)

type GatalogRepository interface {
}

type catalogRepository struct {
	db *gorm.DB
}

func NewCatalogRepository(db *gorm.DB) GatalogRepository {
	return &catalogRepository{
		db: db,
	}
}

package repository

import (
	"errors"
	"go-ecommerce/internal/domain"
	"log"

	"gorm.io/gorm"
)

type GatalogRepository interface {
	CreateCategory(e *domain.Category) error
	FindCategories() ([]*domain.Category, error)
	FindCategoryById(id int) (*domain.Category, error)
	EditCategory(e *domain.Category) (*domain.Category, error)
	DeleteCategory(id int) error
}

type catalogRepository struct {
	db *gorm.DB
}

// CreateCategory implements GatalogRepository.
func (c *catalogRepository) CreateCategory(e *domain.Category) error {
	err := c.db.Create(&e).Error

	if err != nil {
		log.Printf("errir: %v", err)
		return errors.New("create category failed")
	}

	return nil
}

// DeleteCategory implements GatalogRepository.
func (c *catalogRepository) DeleteCategory(id int) error {
	err := c.db.Delete(&domain.Category{}, id).Error

	if err != nil {
		return errors.New("could not delete category")
	}

	return nil
}

// EditCategory implements GatalogRepository.
func (c *catalogRepository) EditCategory(e *domain.Category) (*domain.Category, error) {
	err := c.db.Save(&e).Error

	if err != nil {
		return nil, errors.New("could not create category")
	}

	return e, nil
}

// FindCategories implements GatalogRepository.
func (c *catalogRepository) FindCategories() ([]*domain.Category, error) {
	var categories []*domain.Category
	err := c.db.Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// FindCategoryById implements GatalogRepository.
func (c *catalogRepository) FindCategoryById(id int) (*domain.Category, error) {
	var category *domain.Category

	err := c.db.First(category, id)

	if err != nil {
		return nil, errors.New("Category not found")
	}

	return category, nil
}

func NewCatalogRepository(db *gorm.DB) GatalogRepository {
	return &catalogRepository{
		db: db,
	}
}

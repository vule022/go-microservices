package service

import (
	"errors"
	"go-ecommerce/config"
	"go-ecommerce/internal/domain"
	"go-ecommerce/internal/dto"
	"go-ecommerce/internal/helper"
	"go-ecommerce/internal/repository"
)

type CatalogService struct {
	Repo   repository.GatalogRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s CatalogService) CreateCategory(input dto.CreateCategoryRequest) error {
	err := s.Repo.CreateCategory(&domain.Category{
		Name:         input.Name,
		ImageUrl:     input.ImageUrl,
		DisplayOrder: input.DisplayOrder,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s CatalogService) EditCategory(id int, input dto.CreateCategoryRequest) (*domain.Category, error) {
	cat, err := s.Repo.FindCategoryById(id)

	if err != nil {
		return nil, errors.New("category does not exist")
	}

	if len(input.Name) > 0 {
		cat.Name = input.Name
	}

	if len(input.ImageUrl) > 0 {
		cat.ImageUrl = input.ImageUrl
	}

	if input.ParentId > 0 {
		cat.ParentId = input.ParentId
	}

	if input.DisplayOrder > 0 {
		cat.DisplayOrder = input.DisplayOrder
	}

	updatedCat, err := s.Repo.EditCategory(cat)

	if err != nil {
		return nil, errors.New("cat does not exist")
	}

	return updatedCat, nil
}

func (s CatalogService) DeleteCategory(id int) error {
	err := s.Repo.DeleteCategory(id)

	if err != nil {
		return errors.New("cat does not exist")
	}

	return nil
}

func (s CatalogService) ListCategories() ([]*domain.Category, error) {
	categories, err := s.Repo.FindCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s CatalogService) GetCategory(id int) (*domain.Category, error) {
	cat, err := s.Repo.FindCategoryById(id)

	if err != nil {
		return nil, err
	}

	return cat, nil
}

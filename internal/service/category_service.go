package service

import (
	"kasir-api/internal/domain"
	"kasir-api/internal/repository"
)

// CategoryService handles business logic for categories
type CategoryService struct {
	repo *repository.CategoryRepository
}

// NewCategoryService creates a new CategoryService
func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// GetAll returns all categories
func (s *CategoryService) GetAll() []domain.Category {
	return s.repo.GetAll()
}

// GetByID returns a category by ID
func (s *CategoryService) GetByID(id int) (*domain.Category, error) {
	return s.repo.GetByID(id)
}

// Create adds a new category
func (s *CategoryService) Create(category domain.Category) domain.Category {
	return s.repo.Create(category)
}

// Update modifies an existing category
func (s *CategoryService) Update(id int, category domain.Category) (*domain.Category, error) {
	return s.repo.Update(id, category)
}

// Delete removes a category by ID
func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}

package repository

import (
	"errors"
	"kasir-api/internal/domain"
)

var ErrCategoryNotFound = errors.New("category tidak ditemukan")

// CategoryRepository handles data access for categories
type CategoryRepository struct {
	categories []domain.Category
	nextID     int
}

// NewCategoryRepository creates a new CategoryRepository with initial data
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		categories: []domain.Category{
			{ID: 1, Name: "Makanan", Description: "Produk makanan"},
			{ID: 2, Name: "Minuman", Description: "Produk minuman"},
		},
		nextID: 3,
	}
}

// GetAll returns all categories
func (r *CategoryRepository) GetAll() []domain.Category {
	return r.categories
}

// GetByID returns a category by ID
func (r *CategoryRepository) GetByID(id int) (*domain.Category, error) {
	for _, c := range r.categories {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, ErrCategoryNotFound
}

// Create adds a new category
func (r *CategoryRepository) Create(category domain.Category) domain.Category {
	category.ID = r.nextID
	r.nextID++
	r.categories = append(r.categories, category)
	return category
}

// Update modifies an existing category
func (r *CategoryRepository) Update(id int, category domain.Category) (*domain.Category, error) {
	for i := range r.categories {
		if r.categories[i].ID == id {
			category.ID = id
			r.categories[i] = category
			return &category, nil
		}
	}
	return nil, ErrCategoryNotFound
}

// Delete removes a category by ID
func (r *CategoryRepository) Delete(id int) error {
	for i, c := range r.categories {
		if c.ID == id {
			r.categories = append(r.categories[:i], r.categories[i+1:]...)
			return nil
		}
	}
	return ErrCategoryNotFound
}

package service

import (
	"kasir-api/internal/domain"
	"kasir-api/internal/repository"
)

// ProductService handles business logic for products
type ProductService struct {
	productRepo  *repository.ProductRepository
	categoryRepo *repository.CategoryRepository
}

// NewProductService creates a new ProductService
func NewProductService(productRepo *repository.ProductRepository, categoryRepo *repository.CategoryRepository) *ProductService {
	return &ProductService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

// GetAll returns all products
func (s *ProductService) GetAll() []domain.Product {
	return s.productRepo.GetAll()
}

// GetByID returns a product by ID with category name (JOIN)
func (s *ProductService) GetByID(id int) (*domain.ProductDetail, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// JOIN: Get category name
	categoryName := ""
	category, err := s.categoryRepo.GetByID(product.CategoryID)
	if err == nil {
		categoryName = category.Name
	}

	// Build ProductDetail with category name
	detail := &domain.ProductDetail{
		ID:           product.ID,
		Nama:         product.Nama,
		Harga:        product.Harga,
		Stok:         product.Stok,
		CategoryID:   product.CategoryID,
		CategoryName: categoryName,
	}

	return detail, nil
}

// Create adds a new product
func (s *ProductService) Create(product domain.Product) domain.Product {
	return s.productRepo.Create(product)
}

// Update modifies an existing product
func (s *ProductService) Update(id int, product domain.Product) (*domain.Product, error) {
	return s.productRepo.Update(id, product)
}

// Delete removes a product by ID
func (s *ProductService) Delete(id int) error {
	return s.productRepo.Delete(id)
}

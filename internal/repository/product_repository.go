package repository

import (
	"errors"
	"kasir-api/internal/domain"
)

var ErrProductNotFound = errors.New("produk tidak ditemukan")

// ProductRepository handles data access for products
type ProductRepository struct {
	products []domain.Product
	nextID   int
}

// NewProductRepository creates a new ProductRepository with initial data
func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: []domain.Product{
			{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10, CategoryID: 1},
			{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40, CategoryID: 2},
			{ID: 3, Nama: "kecap", Harga: 12000, Stok: 20, CategoryID: 1},
		},
		nextID: 4,
	}
}

// GetAll returns all products
func (r *ProductRepository) GetAll() []domain.Product {
	return r.products
}

// GetByID returns a product by ID
func (r *ProductRepository) GetByID(id int) (*domain.Product, error) {
	for _, p := range r.products {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, ErrProductNotFound
}

// Create adds a new product
func (r *ProductRepository) Create(product domain.Product) domain.Product {
	product.ID = r.nextID
	r.nextID++
	r.products = append(r.products, product)
	return product
}

// Update modifies an existing product
func (r *ProductRepository) Update(id int, product domain.Product) (*domain.Product, error) {
	for i := range r.products {
		if r.products[i].ID == id {
			product.ID = id
			r.products[i] = product
			return &product, nil
		}
	}
	return nil, ErrProductNotFound
}

// Delete removes a product by ID
func (r *ProductRepository) Delete(id int) error {
	for i, p := range r.products {
		if p.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return ErrProductNotFound
}

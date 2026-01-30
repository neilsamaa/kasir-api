package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kasir-api/internal/handler"
	"kasir-api/internal/repository"
	"kasir-api/internal/service"
)

func main() {
	// Initialize repositories (Data Layer)
	categoryRepo := repository.NewCategoryRepository()
	productRepo := repository.NewProductRepository()

	// Initialize services (Domain/Business Logic Layer)
	categoryService := service.NewCategoryService(categoryRepo)
	productService := service.NewProductService(productRepo, categoryRepo)

	// Initialize handlers (Presentation Layer)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewProductHandler(productService)

	// Register routes

	// Category routes
	// GET /categories - Get all categories
	// POST /categories - Create category
	http.HandleFunc("/categories", categoryHandler.HandleCategories)

	// GET /categories/{id} - Get category by ID
	// PUT /categories/{id} - Update category
	// DELETE /categories/{id} - Delete category
	http.HandleFunc("/categories/", categoryHandler.HandleCategoryByID)

	// Product routes
	// GET /api/produk - Get all products
	// POST /api/produk - Create product
	http.HandleFunc("/api/produk", productHandler.HandleProducts)

	// GET /api/produk/{id} - Get product detail WITH category_name
	// PUT /api/produk/{id} - Update product
	// DELETE /api/produk/{id} - Delete product
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	fmt.Println("Server running di localhost:8080")
	fmt.Println("Layered Architecture:")
	fmt.Println("  - Domain:     internal/domain/")
	fmt.Println("  - Repository: internal/repository/")
	fmt.Println("  - Service:    internal/service/")
	fmt.Println("  - Handler:    internal/handler/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server:", err)
	}
}

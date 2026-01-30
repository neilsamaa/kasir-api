package domain

// Product represents a product entity
type Product struct {
	ID         int    `json:"id"`
	Nama       string `json:"nama"`
	Harga      int    `json:"harga"`
	Stok       int    `json:"stok"`
	CategoryID int    `json:"category_id"`
}

// ProductDetail is used for GET detail response with category name
type ProductDetail struct {
	ID           int    `json:"id"`
	Nama         string `json:"nama"`
	Harga        int    `json:"harga"`
	Stok         int    `json:"stok"`
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

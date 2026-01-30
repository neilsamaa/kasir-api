# Kasir API

REST API untuk aplikasi kasir, dibangun dengan **Go** menggunakan **Layered Architecture**.

## 🏗️ Arsitektur

Project ini menggunakan Clean Architecture dengan pemisahan layer:

```
kasir-api/
├── cmd/api/main.go           # Application entry point
├── internal/
│   ├── domain/               # Entities (innermost layer)
│   │   ├── category.go
│   │   └── product.go
│   ├── repository/           # Data access layer
│   │   ├── category_repository.go
│   │   └── product_repository.go
│   ├── service/              # Business logic / use cases
│   │   ├── category_service.go
│   │   └── product_service.go
│   └── handler/              # HTTP handlers (presentation layer)
│       ├── category_handler.go
│       └── product_handler.go
└── go.mod
```

## 🚀 Cara Menjalankan

```bash
# Development
go run ./cmd/api

# Build & Run
go build -o kasir-api.exe ./cmd/api
./kasir-api.exe
```

Server berjalan di `http://localhost:8080`

## 📡 API Endpoints

### Health Check
| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/health` | Cek status API |

### Categories
| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/categories` | Ambil semua kategori |
| POST | `/categories` | Tambah kategori baru |
| GET | `/categories/{id}` | Ambil detail kategori |
| PUT | `/categories/{id}` | Update kategori |
| DELETE | `/categories/{id}` | Hapus kategori |

### Products
| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/produk` | Ambil semua produk |
| POST | `/api/produk` | Tambah produk baru |
| GET | `/api/produk/{id}` | Ambil detail produk (dengan category_name) |
| PUT | `/api/produk/{id}` | Update produk |
| DELETE | `/api/produk/{id}` | Hapus produk |

## 📝 Contoh Request

### Tambah Category
```bash
curl -X POST http://localhost:8080/categories \
  -H "Content-Type: application/json" \
  -d '{"name":"Snack","description":"Makanan ringan"}'
```

### Tambah Product
```bash
curl -X POST http://localhost:8080/api/produk \
  -H "Content-Type: application/json" \
  -d '{"nama":"Chitato","harga":15000,"stok":50,"category_id":1}'
```

### Get Product Detail (dengan JOIN category)
```bash
curl http://localhost:8080/api/produk/1
```

Response:
```json
{
  "id": 1,
  "nama": "Indomie Godog",
  "harga": 3500,
  "stok": 10,
  "category_id": 1,
  "category_name": "Makanan"
}
```

## 📚 Referensi Arsitektur

- [Clean Architecture - Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Presentation Domain Data Layering - Martin Fowler](https://martinfowler.com/bliki/PresentationDomainDataLayering.html)
- [Go Project Layout](https://github.com/golang-standards/project-layout)

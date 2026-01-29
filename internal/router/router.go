package router

import (
	"kasir-api/internal/handler"
	"kasir-api/internal/repository"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(pool *pgxpool.Pool) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", handler.CheckHealth)

	queries := repository.New(pool)
	registerProductsRoutes(mux, queries)
	registerCategoriesRoutes(mux, queries)

	return mux
}

func registerProductsRoutes(mux *http.ServeMux, queries *repository.Queries) {
	productHandler := handler.NewProductHandler(queries)

	mux.HandleFunc("GET /api/products", productHandler.GetAllProducts)
	mux.HandleFunc("POST /api/products", handler.AddProduct)
	mux.HandleFunc("GET /api/products/{id}", handler.GetProductByID)
	mux.HandleFunc("PUT /api/products/{id}", handler.UpdateProductByID)
	mux.HandleFunc("DELETE /api/products/{id}", handler.DeleteProductByID)
}

func registerCategoriesRoutes(mux *http.ServeMux, queries *repository.Queries) {
	mux.HandleFunc("GET /api/categories", handler.GetAllCategories)
	mux.HandleFunc("POST /api/categories", handler.AddCategory)
	mux.HandleFunc("GET /api/categories/{id}", handler.GetCategoryByID)
	mux.HandleFunc("PUT /api/categories/{id}", handler.UpdateCategoryByID)
	mux.HandleFunc("DELETE /api/categories/{id}", handler.DeleteCategoryByID)
}

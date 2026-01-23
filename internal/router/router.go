package router

import (
	"kasir-api/internal/controllers"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", controllers.CheckHealth)

	mux.HandleFunc("GET /api/product", controllers.GetAllProducts)
	mux.HandleFunc("POST /api/product", controllers.AddProduct)
	mux.HandleFunc("GET /api/product/{id}", controllers.GetProductByID)
	mux.HandleFunc("PUT /api/product/{id}", controllers.UpdateProductByID)
	mux.HandleFunc("DELETE /api/product/{id}", controllers.DeleteProductByID)

	mux.HandleFunc("GET /api/category", controllers.GetAllCategories)
	mux.HandleFunc("POST /api/category", controllers.AddCategory)
	mux.HandleFunc("GET /api/category/{id}", controllers.GetCategoryByID)
	mux.HandleFunc("PUT /api/category/{id}", controllers.UpdateCategoryByID)
	mux.HandleFunc("DELETE /api/category/{id}", controllers.DeleteCategoryByID)

	return mux
}

package router

import (
	"kasir-api/internal/controllers"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", controllers.CheckHealth)

	registerProductsRoutes(mux)
	registerCategoriesRoutes(mux)

	return mux
}

func registerProductsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/products", controllers.GetAllProducts)
	mux.HandleFunc("POST /api/products", controllers.AddProduct)
	mux.HandleFunc("GET /api/products/{id}", controllers.GetProductByID)
	mux.HandleFunc("PUT /api/products/{id}", controllers.UpdateProductByID)
	mux.HandleFunc("DELETE /api/products/{id}", controllers.DeleteProductByID)
}

func registerCategoriesRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/categories", controllers.GetAllCategories)
	mux.HandleFunc("POST /api/categories", controllers.AddCategory)
	mux.HandleFunc("GET /api/categories/{id}", controllers.GetCategoryByID)
	mux.HandleFunc("PUT /api/categories/{id}", controllers.UpdateCategoryByID)
	mux.HandleFunc("DELETE /api/categories/{id}", controllers.DeleteCategoryByID)
}

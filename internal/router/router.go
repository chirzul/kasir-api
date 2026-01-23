package router

import (
	"kasir-api/internal/controllers"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", controllers.CheckHealth)

	mux.HandleFunc("GET /api/product", controllers.GetAllProduct)
	mux.HandleFunc("POST /api/product", controllers.AddProduct)
	mux.HandleFunc("GET /api/product/{id}", controllers.GetProductByID)
	mux.HandleFunc("PUT /api/product/{id}", controllers.UpdateProductByID)
	mux.HandleFunc("DELETE /api/product/{id}", controllers.DeleteProductByID)

	return mux
}

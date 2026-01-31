package router

import (
	"kasir-api/internal/handler"

	"github.com/gofiber/fiber/v3"
)

type Handlers struct {
	ProductHandler  *handler.ProductHandler
	CategoryHandler *handler.CategoryHandler
}

func RegisterRoutes(app fiber.Router, h *Handlers) {
	api := app.Group("/api")

	registerProductsRoutes(api, h.ProductHandler)
	registerCategoriesRoutes(api, h.CategoryHandler)
}

func registerProductsRoutes(api fiber.Router, h *handler.ProductHandler) {
	v1 := api.Group("/v1")
	product := v1.Group("/products")
	product.Get("/", h.GetAllProducts)
	product.Post("/", h.AddProduct)
	product.Get("/:id", h.GetProductByID)
	product.Put("/:id", h.UpdateProductByID)
	product.Delete("/:id", h.DeleteProductByID)
}

func registerCategoriesRoutes(api fiber.Router, h *handler.CategoryHandler) {
	v1 := api.Group("/v1")
	categories := v1.Group("/categories")
	categories.Get("/", h.GetAllCategories)
	categories.Post("/", h.AddCategory)
	categories.Get("/:id", h.GetCategoryByID)
	categories.Put("/:id", h.UpdateCategoryByID)
	categories.Delete("/:id", h.DeleteCategoryByID)
}

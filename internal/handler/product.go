package handler

import (
	"errors"
	"kasir-api/internal/dto"
	"kasir-api/internal/pkg/utils"
	"kasir-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.service.GetAllProducts(c.Context())
	if errors.Is(err, utils.ErrNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"statusCode": fiber.StatusNotFound,
			"statusDesc": "NOT_FOUND",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": fiber.StatusInternalServerError,
			"statusDesc": "INTERNAL_SERVER_ERROR",
			"error":      err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"statusDesc": "OK",
		"data":       products,
	})
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := h.service.GetProductById(c.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrNotFound):
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"statusCode": fiber.StatusNotFound,
				"statusDesc": "NOT_FOUND",
				"error":      err.Error(),
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"statusDesc": "INTERNAL_SERVER_ERROR",
				"error":      err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"statusDesc": "OK",
		"data":       product,
	})
}

func (h *ProductHandler) AddProduct(c *fiber.Ctx) error {
	var req dto.AddProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"statusDesc": "BAD_REQUEST",
			"error":      err.Error(),
		})
	}

	if err := h.service.AddProduct(c.Context(), req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": fiber.StatusInternalServerError,
			"statusDesc": "INTERNAL_SERVER_ERROR",
			"error":      err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"statusCode": fiber.StatusCreated,
		"statusDesc": "CREATED",
	})
}

func (h *ProductHandler) UpdateProductByID(c *fiber.Ctx) error {
	var req dto.UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"statusDesc": "BAD_REQUEST",
			"error":      err.Error(),
		})
	}

	req.ID = c.Params("id")
	if err := h.service.UpdateProductById(c.Context(), req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"statusCode": fiber.StatusInternalServerError,
			"statusDesc": "INTERNAL_SERVER_ERROR",
			"error":      err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"statusDesc": "OK",
	})
}

func (h *ProductHandler) DeleteProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.service.DeleteProductById(c.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"statusCode": fiber.StatusNotFound,
				"statusDesc": "NOT_FOUND",
				"error":      err.Error(),
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"statusDesc": "INTERNAL_SERVER_ERROR",
				"error":      err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"statusCode": fiber.StatusOK,
		"statusDesc": "OK",
	})
}

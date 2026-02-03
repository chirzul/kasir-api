package handler

import (
	"errors"
	"kasir-api/internal/dto"
	"kasir-api/internal/pkg/utils"
	"kasir-api/internal/service"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

// GetAllProducts godoc
//
//	@Summary		List products
//	@Description	get list products
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.SuccessResponse{data=[]dto.ProductResponse}
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/products [get]
func (h *ProductHandler) GetAllProducts(c fiber.Ctx) error {
	products, err := h.service.GetAllProducts(c.RequestCtx())
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

// GetProductByID godoc
//
//	@Summary		Detail product
//	@Description	get detail product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Product ID"
//	@Success		200	{object}	dto.SuccessResponse{data=dto.ProductResponse}
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/products/{id} [get]
func (h *ProductHandler) GetProductByID(c fiber.Ctx) error {
	id := c.Params("id")

	product, err := h.service.GetProductById(c.RequestCtx(), id)
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

// AddProduct godoc
//
//	@Summary		Add product
//	@Description	add new product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			req	body		dto.AddProductRequest	true	"Product Request"
//	@Success		201	{object}	dto.SuccessResponse{data=nil}
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/products [post]
func (h *ProductHandler) AddProduct(c fiber.Ctx) error {
	var req dto.AddProductRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"statusDesc": "BAD_REQUEST",
			"error":      err.Error(),
		})
	}

	if err := h.service.AddProduct(c.RequestCtx(), req); err != nil {
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

// UpdateProductById godoc
//
//	@Summary		Update product
//	@Description	update product data
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			req	body		dto.UpdateProductRequest	true	"Product Request"
//	@Param			id	path		int							true	"Product ID"
//	@Success		201	{object}	dto.SuccessResponse{data=nil}
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/products/{id} [put]
func (h *ProductHandler) UpdateProductByID(c fiber.Ctx) error {
	var req dto.UpdateProductRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"statusDesc": "BAD_REQUEST",
			"error":      err.Error(),
		})
	}

	req.ID = c.Params("id")
	if err := h.service.UpdateProductById(c.RequestCtx(), req); err != nil {
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

// DeleteProductByID godoc
//
//	@Summary		Delete product
//	@Description	delete product by id
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Product ID"
//	@Success		200	{object}	dto.SuccessResponse{data=nil}
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/products/{id} [delete]
func (h *ProductHandler) DeleteProductByID(c fiber.Ctx) error {
	id := c.Params("id")

	err := h.service.DeleteProductById(c.RequestCtx(), id)
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

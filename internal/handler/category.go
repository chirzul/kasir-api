package handler

import (
	"errors"
	"kasir-api/internal/dto"
	"kasir-api/internal/pkg/utils"
	"kasir-api/internal/service"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(s service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

// GetAllCategories godoc
//
//	@Summary		List categories
//	@Description	get list categories
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.SuccessResponse{data=[]dto.CategoryResponse}
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/categories [get]
func (h *CategoryHandler) GetAllCategories(c fiber.Ctx) error {
	categories, err := h.service.GetAllCategories(c.RequestCtx())
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
		"data":       categories,
	})
}

// GetCategoryByID godoc
//
//	@Summary		Detail category
//	@Description	get detail category
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Category ID"
//	@Success		200	{object}	dto.SuccessResponse{data=dto.CategoryResponse}
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(c fiber.Ctx) error {
	id := c.Params("id")

	category, err := h.service.GetCategoryById(c.RequestCtx(), id)
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
		"data":       category,
	})
}

// AddCategory godoc
//
//	@Summary		Add category
//	@Description	add new category
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			req	body		dto.AddCategoryRequest	true	"Category Request"
//	@Success		201	{object}	dto.SuccessResponse{data=nil}
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/categories [post]
func (h *CategoryHandler) AddCategory(c fiber.Ctx) error {
	var req dto.AddCategoryRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"statusDesc": "BAD_REQUEST",
			"error":      err.Error(),
		})
	}

	if err := h.service.AddCategory(c.RequestCtx(), req); err != nil {
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

// UpdateCategoryById godoc
//
//	@Summary		Update category
//	@Description	update category data
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			req	body		dto.UpdateCategoryRequest	true	"Category Request"
//	@Param			id	path		string						true	"Category ID"
//	@Success		201	{object}	dto.SuccessResponse{data=nil}
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/categories/{id} [put]
func (h *CategoryHandler) UpdateCategoryByID(c fiber.Ctx) error {
	var req dto.UpdateCategoryRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"statusCode": fiber.StatusBadRequest,
			"statusDesc": "BAD_REQUEST",
			"error":      err.Error(),
		})
	}

	req.ID = c.Params("id")
	if err := h.service.UpdateCategoryById(c.RequestCtx(), req); err != nil {
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

// DeleteCategoryByID godoc
//
//	@Summary		Delete category
//	@Description	delete category by id
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Category ID"
//	@Success		200	{object}	dto.SuccessResponse{data=nil}
//	@Failure		404	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/categories/{id} [delete]
func (h *CategoryHandler) DeleteCategoryByID(c fiber.Ctx) error {
	id := c.Params("id")

	err := h.service.DeleteCategoryById(c.RequestCtx(), id)
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

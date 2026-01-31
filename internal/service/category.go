package service

import (
	"context"
	"kasir-api/internal/dto"
	"kasir-api/internal/pkg/utils"
	"kasir-api/internal/repository"

	"github.com/jackc/pgx/v5/pgtype"
)

type CategoryService interface {
	GetAllCategories(ctx context.Context) ([]dto.CategoryResponse, error)
	GetCategoryById(ctx context.Context, id string) (*dto.CategoryResponse, error)
	AddCategory(ctx context.Context, req dto.AddCategoryRequest) error
	UpdateCategoryById(ctx context.Context, req dto.UpdateCategoryRequest) error
	DeleteCategoryById(ctx context.Context, id string) error
}

type categoryService struct {
	repository *repository.Queries
}

func NewCategoryService(r *repository.Queries) CategoryService {
	return &categoryService{repository: r}
}

func (s *categoryService) GetAllCategories(ctx context.Context) ([]dto.CategoryResponse, error) {
	var resp []dto.CategoryResponse
	categories, err := s.repository.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, utils.ErrNotFound
	}

	for _, c := range categories {
		resp = append(resp, dto.CategoryResponse{
			ID:          c.ID,
			Name:        c.Name,
			Description: c.Description.String,
		})
	}

	return resp, nil
}

func (s *categoryService) GetCategoryById(ctx context.Context, id string) (*dto.CategoryResponse, error) {
	category, err := s.repository.GetCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description.String,
	}, nil
}

func (s *categoryService) AddCategory(ctx context.Context, req dto.AddCategoryRequest) error {
	var desc pgtype.Text
	desc.Scan(req.Description)
	return s.repository.AddCategory(ctx, repository.AddCategoryParams{
		ID:          req.ID,
		Name:        req.Name,
		Description: desc,
	})
}

func (s *categoryService) UpdateCategoryById(ctx context.Context, req dto.UpdateCategoryRequest) error {
	var desc pgtype.Text
	desc.Scan(req.Description)
	rows, err := s.repository.UpdateCategoryById(ctx, repository.UpdateCategoryByIdParams{
		ID:          req.ID,
		Name:        req.Name,
		Description: desc,
	})
	if err != nil {
		return err
	}
	if rows.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}

func (s *categoryService) DeleteCategoryById(ctx context.Context, id string) error {
	rows, err := s.repository.DeleteCategoryById(ctx, id)
	if err != nil {
		return err
	}
	if rows.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}

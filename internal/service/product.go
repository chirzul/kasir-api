package service

import (
	"context"
	"fmt"
	"kasir-api/internal/dto"
	"kasir-api/internal/pkg/utils"
	"kasir-api/internal/repository"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spf13/cast"
)

type ProductService interface {
	GetAllProducts(ctx context.Context) ([]dto.ProductResponse, error)
	GetProductById(ctx context.Context, id string) (*dto.ProductResponse, error)
	AddProduct(ctx context.Context, req dto.AddProductRequest) error
	UpdateProductById(ctx context.Context, req dto.UpdateProductRequest) error
	DeleteProductById(ctx context.Context, id string) error
}

type productService struct {
	repository *repository.Queries
}

func NewProductService(r *repository.Queries) ProductService {
	return &productService{repository: r}
}

func (s *productService) GetAllProducts(ctx context.Context) ([]dto.ProductResponse, error) {
	var resp []dto.ProductResponse
	products, err := s.repository.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, utils.ErrNotFound
	}

	for _, p := range products {
		price, _ := p.Price.Float64Value()
		resp = append(resp, dto.ProductResponse{
			ID:           p.ID,
			Name:         p.Name,
			Price:        price.Float64,
			Stock:        int(p.Stock),
			CategoryName: p.CategoryName,
			CategoryDesc: p.CategoryDescription.String,
		})
	}

	return resp, nil
}

func (s *productService) GetProductById(ctx context.Context, id string) (*dto.ProductResponse, error) {
	product, err := s.repository.GetProductById(ctx, cast.ToInt64(id))
	if err != nil {
		return nil, err
	}

	price, _ := product.Price.Float64Value()
	return &dto.ProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Price:        price.Float64,
		Stock:        int(product.Stock),
		CategoryName: product.CategoryName,
		CategoryDesc: product.CategoryDescription.String,
	}, nil
}

func (s *productService) AddProduct(ctx context.Context, req dto.AddProductRequest) error {
	var price pgtype.Numeric
	price.Scan(fmt.Sprintf("%.2f", req.Price))
	return s.repository.AddProduct(ctx, repository.AddProductParams{
		Name:       req.Name,
		Price:      price,
		Stock:      int16(req.Stock),
		CategoryID: req.CategoryId,
	})
}

func (s *productService) UpdateProductById(ctx context.Context, req dto.UpdateProductRequest) error {
	var price pgtype.Numeric
	price.Scan(fmt.Sprintf("%.2f", req.Price))
	rows, err := s.repository.UpdateProductById(ctx, repository.UpdateProductByIdParams{
		Name:       req.Name,
		Price:      price,
		Stock:      int16(req.Stock),
		CategoryID: req.CategoryId,
		ID:         cast.ToInt64(req.ID),
	})
	if err != nil {
		return err
	}
	if rows.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}

func (s *productService) DeleteProductById(ctx context.Context, id string) error {
	rows, err := s.repository.DeleteProductById(ctx, cast.ToInt64(id))
	if err != nil {
		return err
	}
	if rows.RowsAffected() == 0 {
		return utils.ErrNotFound
	}
	return nil
}

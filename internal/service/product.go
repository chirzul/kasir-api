package service

import (
	"context"
	"kasir-api/internal/repository"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spf13/cast"
)

type ProductService interface {
	GetAllProducts(ctx context.Context) ([]repository.GetAllProductsRow, error)
	GetProductById(ctx context.Context, id string) (repository.GetProductByIdRow, error)
	AddProduct(ctx context.Context) error
	UpdateProductById(ctx context.Context) error
	DeleteProductById(ctx context.Context, id int64) error
}

type productService struct {
	q *repository.Queries
}

func NewProductService(q *repository.Queries) ProductService {
	return &productService{q: q}
}

func (s *productService) GetAllProducts(ctx context.Context) ([]repository.GetAllProductsRow, error) {
	return s.q.GetAllProducts(ctx)
}

func (s *productService) GetProductById(ctx context.Context, id string) (repository.GetProductByIdRow, error) {
	return s.q.GetProductById(ctx, cast.ToInt64(id))
}

func (s *productService) AddProduct(ctx context.Context) error {
	return s.q.AddProduct(ctx, repository.AddProductParams{
		Name:       "",
		Price:      pgtype.Numeric{},
		Stock:      2,
		CategoryID: "",
	})
}

func (s *productService) UpdateProductById(ctx context.Context) error {
	return s.q.UpdateProductById(ctx, repository.UpdateProductByIdParams{
		ID:    0,
		Name:  pgtype.Text{},
		Price: pgtype.Numeric{},
		Stock: pgtype.Int2{},
	})
}

func (s *productService) DeleteProductById(ctx context.Context, id int64) error {
	return s.q.DeleteProductById(ctx, id)
}

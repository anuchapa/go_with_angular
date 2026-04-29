package products

import (
	"context"
	"goBackend/internal/adapter/db/model"
	"goBackend/internal/products/dtos"
	"log"
)

type Service interface {
	GetAllProducts(ctx context.Context) (*dtos.ProductListResponse, error)
	GetProductByID(ctx context.Context, id int64) (*dtos.ProductResponse, error)
	CreateProduct(ctx context.Context, product *dtos.ProductListCreateRequest) (*dtos.ProductListResponse, error)
	DeleteProduct(ctx context.Context, id int64) (int64,error)
}

type serviceImp struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImp{repo: repo}
}

func (s *serviceImp) GetAllProducts(ctx context.Context) (*dtos.ProductListResponse, error) {
	products, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*dtos.ProductResponse, len(products))

	for i, p := range products {
		items[i] = &dtos.ProductResponse{
			ID:          p.ID,
			ProductCode: p.ProductCode,
		}
	}

	resp := &dtos.ProductListResponse{
		Data: items,
	}
	return resp, nil
}
func (s *serviceImp) GetProductByID(ctx context.Context, id int64) (*dtos.ProductResponse, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil,nil
	}

	resp := &dtos.ProductResponse{
		ID:          product.ID,
		ProductCode: product.ProductCode,
	}
	return resp, nil
}

func (s *serviceImp) CreateProduct(ctx context.Context, productBody *dtos.ProductListCreateRequest) (*dtos.ProductListResponse, error) {
	products := make([]*model.Product, len(*productBody))
	for i, p := range *productBody {
		products[i] = &model.Product{
			ProductCode: p.ProductCode,
		}
	}

	if err := s.repo.Create(ctx, products); err != nil {
		return nil, err
	}
	for _, p := range products {
		log.Printf("%d", p.ID)
	}

	items := make([]*dtos.ProductResponse, len(products))

	for i, p := range products {
		items[i] = &dtos.ProductResponse{
			ID:          p.ID,
			ProductCode: p.ProductCode,
		}
	}

	resp := &dtos.ProductListResponse{
		Data: items,
	}

	return resp, nil
}

func (s *serviceImp) DeleteProduct(ctx context.Context, id int64) (int64,error) {
	rowsAffected, err := s.repo.Delete(ctx, id)
	if err != nil {
		return 0,err
	}
	return rowsAffected,nil
}

package product

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateProduct(ctx context.Context, name, description string, price float64) (*Product, error) {
	p := &Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	err := s.repo.Create(ctx, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *Service) GetProducts(ctx context.Context) ([]Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetProductByID(ctx context.Context, id int64) (*Product, error) {
	return s.repo.GetByID(ctx, id)
}

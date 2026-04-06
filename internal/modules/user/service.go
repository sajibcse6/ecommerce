package user

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, name, email string) (*User, error) {
	user := &User{
		Name:  name,
		Email: email,
	}

	err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetUsers(ctx context.Context) ([]User, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) GetUserByID(ctx context.Context, id int64) (*User, error) {
	return s.repo.GetByID(ctx, id)
}

package auth

import (
	"context"
	"errors"

	"ecommerce/internal/modules/user"
	"ecommerce/internal/pkg/jwt"
)

type Service struct {
	userRepo *user.Repository
	secret   string
}

func NewService(userRepo *user.Repository, secret string) *Service {
	return &Service{
		userRepo: userRepo,
		secret:   secret,
	}
}

func (s *Service) Register(ctx context.Context, name, email string) (*user.User, error) {
	u := &user.User{
		Name:  name,
		Email: email,
	}

	err := s.userRepo.Create(ctx, u)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) Login(ctx context.Context, email string) (string, error) {
	// Simple login (no password yet for learning)

	users, err := s.userRepo.GetAll(ctx)

	if err != nil {
		return "", err
	}

	for _, u := range users {
		if u.Email == email {
			return jwt.GenerateToken(u.ID, s.secret)
		}
	}

	return "", errors.New("Invalid credentials")
}

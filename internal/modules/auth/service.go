package auth

import (
	"context"
	"errors"

	"ecommerce/internal/modules/user"
	"ecommerce/internal/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
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

func (s *Service) Register(ctx context.Context, name, email, password string) (*user.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err !=  nil {
		return nil, err
	}

	u := &user.User{
		Name:  name,
		Email: email,
		Password: string(hashed),
	}

	err = s.userRepo.Create(ctx, u)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	return jwt.GenerateToken(user.ID, s.secret)
}

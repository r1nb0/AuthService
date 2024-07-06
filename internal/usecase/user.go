package usecase

import (
	"AuthService/internal/domain"
	"context"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) UserUseCase {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetAll(ctx context.Context) ([]*domain.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *userService) Update(ctx context.Context, id int, user *domain.UserDTO) error {
	return s.repo.Update(ctx, id, user)
}

func (s *userService) GetByID(ctx context.Context, id int) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

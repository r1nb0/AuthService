package usecase

import (
	"context"
	"github.com/r1nb0/UserService/internal/domain"
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

func (s *userService) Update(ctx context.Context, id int, dto *domain.UserDTO) error {
	return s.repo.Update(ctx, id, dto)
}

func (s *userService) GetByID(ctx context.Context, id int) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

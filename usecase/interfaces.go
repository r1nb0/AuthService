package usecase

import (
	"context"
	"github.com/r1nb0/UserService/domain"
)

type AuthUseCase interface {
	SignIn(ctx context.Context, dto *domain.UserAuthDTO) (string, error)
	SignUp(ctx context.Context, dto *domain.UserDTO) (int, error)
}

type UserUseCase interface {
	GetAll(ctx context.Context) ([]*domain.User, error)
	GetByID(ctx context.Context, id int) (*domain.User, error)
	Update(ctx context.Context, id int, user *domain.UserDTO) error
}

package usecase

import (
	"AuthService/internal/domain"
	"context"
)

type AuthUseCase interface {
	SignIn(ctx context.Context, dto *domain.UserAuthDTO) (string, error)
	SignUp(ctx context.Context, dto *domain.UserDTO) (int, error)
}

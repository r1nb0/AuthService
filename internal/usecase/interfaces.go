package usecase

import (
	"context"
	"github.com/r1nb0/UserService/internal/domain"
)

type UserUseCase interface {
	SignIn(ctx context.Context, dto *domain.AuthenticateUser) (string, error)
	SignUp(ctx context.Context, dto *domain.CreateUser) (int, error)
	GetAll(ctx context.Context) ([]*domain.User, error)
	GetByID(ctx context.Context, id int) (*domain.User, error)
	Update(ctx context.Context, id int, user *domain.UpdateUserGeneralInfo) error
	UpdatePassword(ctx context.Context, id int, password string) error
	UpdateEmail(ctx context.Context, id int, email string) error
}

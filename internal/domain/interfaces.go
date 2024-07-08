package domain

import "context"

type UserRepository interface {
	Create(ctx context.Context, dto *CreateUser) (int, error)
	GetByAuthData(ctx context.Context, dto *AuthenticateUser) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	GetByID(ctx context.Context, id int) (*User, error)
	ExistsByNickname(ctx context.Context, nickname string) bool
	ExistsByEmail(ctx context.Context, email string) bool
	Update(ctx context.Context, id int, dto *UpdateUserGeneralInfo) error
	UpdatePassword(ctx context.Context, id int, password string) error
	UpdateEmail(ctx context.Context, id int, email string) error
}

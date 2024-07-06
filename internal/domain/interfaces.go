package domain

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context, dto *UserDTO) (int, error)
	GetUser(ctx context.Context, dto *UserAuthDTO) (*User, error)
}
